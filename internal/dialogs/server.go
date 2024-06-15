package dialogs

import (
	"context"
	"fmt"
	"net"
	"runtime/debug"

	"github.com/davecgh/go-spew/spew"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	logger "github.com/sirupsen/logrus"
	"github.com/xpoh/otus-work-dialogs/internal/clickhouse"
	"github.com/xpoh/otus-work-dialogs/pkg/grpc/dialogs/v1"
	"google.golang.org/grpc"
	grpcCodes "google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
)

type Config interface {
	GetDialogsURI() string
	GetShardsCount() int
}

type Server struct {
	cfg   Config
	click *clickhouse.Client

	grpcServer *grpc.Server
	dialogs.UnimplementedDialogServiceServer
}

func NewServer(cfg Config, click *clickhouse.Client) *Server {
	dialogsServer := &Server{cfg: cfg, click: click}

	serverOptions := setupServerOptions()
	dialogsServer.grpcServer = grpc.NewServer(serverOptions...)

	dialogs.RegisterDialogServiceServer(dialogsServer.grpcServer, dialogsServer)

	return dialogsServer
}

func (s *Server) Start() {
	grpcBindAddress := s.cfg.GetDialogsURI()

	listener, err := net.Listen("tcp", grpcBindAddress)
	if err != nil {
		logger.Panicf("unable to create grpc listener: %v", err)
	}

	logger.Infof("gRPC server is listening on: %q", grpcBindAddress)

	if err = s.grpcServer.Serve(listener); err != nil {
		logger.Panicf("unable to start server: %v", err)
	}
}

func (s *Server) Stop(_ context.Context) {
	s.grpcServer.GracefulStop()
}

func setupServerOptions() []grpc.ServerOption {
	recoveryOptions := []grpcRecovery.Option{
		grpcRecovery.WithRecoveryHandler(recoveryFunc),
	}

	loggerEntry := logger.NewEntry(logger.StandardLogger())
	loggerOptions := []grpcLogrus.Option{
		grpcLogrus.WithLevels(func(code grpcCodes.Code) logger.Level {
			if code == grpcCodes.OK {
				return logger.DebugLevel
			}

			return logger.ErrorLevel
		}),
	}

	return []grpc.ServerOption{
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcLogrus.UnaryServerInterceptor(loggerEntry, loggerOptions...),
			grpcRecovery.UnaryServerInterceptor(recoveryOptions...),
		)),
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcLogrus.StreamServerInterceptor(loggerEntry, loggerOptions...),
			grpcRecovery.StreamServerInterceptor(recoveryOptions...),
		)),
	}
}

func recoveryFunc(p interface{}) error {
	logger.Errorf("%s\nstacktrace from panic: %v\n", spew.Sprintf("%#v", p), string(debug.Stack()))

	return grpcStatus.Errorf(grpcCodes.Internal, "internal error")
}

func (s *Server) dialogSend(userFrom, userTo, text string) error {
	c := s.click.GetConn()

	shardKey := int(userFrom[0]+userTo[0]) % s.cfg.GetShardsCount()

	if _, err := c.Exec(
		`INSERT INTO "DialogMessage" VALUES ($1,$2,$3,$4)`,
		shardKey,
		userFrom,
		userTo,
		text,
	); err != nil {
		return err
	}

	return nil
}

func (s *Server) Send(ctx context.Context, request *dialogs.SendRequest) (*dialogs.SendResponse, error) {
	if err := s.dialogSend(request.UserFrom, request.UserTo, request.Text); err != nil {
		return nil, err
	}

	return &dialogs.SendResponse{}, nil
}

func (s *Server) dialogList(userFrom, userTo string) ([]*DialogMessage, error) {
	c := s.click.GetConn()

	rows, err := c.Query(
		`select from_user_id, to_user_id, text from "DialogMessage"
          where from_user_id=$1 AND to_user_id=$2`, userFrom, userTo)
	if err != nil {
		return nil, fmt.Errorf("error get messages: %w", err)
	}

	defer rows.Close()

	messages := make([]*DialogMessage, 0)

	for rows.Next() {
		var message DialogMessage

		if err = rows.Scan(
			&message.From,
			&message.To,
			&message.Text,
		); err != nil {
			return nil, fmt.Errorf("error list query: %w", err)
		}

		messages = append(messages, &message)
	}

	return messages, nil
}

func (s *Server) List(ctx context.Context, request *dialogs.ListRequest) (*dialogs.ListResponse, error) {
	messages, err := s.dialogList(request.UserFrom, request.UserTo)
	if err != nil {
		return nil, err
	}

	response := dialogs.ListResponse{}

	for _, message := range messages {
		response.Text = append(response.Text, message.Text)
	}

	return &response, nil
}

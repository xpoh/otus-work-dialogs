package config

import (
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/davecgh/go-spew/spew"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

const (
	envConfigFileName = ".env"
)

type Config struct {
	env *EnvSetting
}

type EnvSetting struct {
	Host string `env:"HOST" env-default:"0.0.0.0" env-description:"Host"`
	Port string `env:"HTTP_PORT" env-default:"8000" env-description:"Http port"`

	ClickhouseHost string `env:"CLICKHOUSE_HOST" env-default:"localhost" env-description:"hostname or IP of clickhouse server"` //nolint:lll
	ClickhousePort uint16 `env:"CLICKHOUSE_PORT" env-default:"80" env-description:"clickhouse instance port"`
	ClickhouseUser string `env:"CLICKHOUSE_USER" env-default:"user" env-description:"clickhouse instance port"`
	MigrationsPath string `env:"MIGRATIONS_PATH" env-default:"scripts/clickhouse" env-description:"path where clickhouse migrations stored"` //nolint:lll
	ShardCount     int    `env:"SHARD_COUNT" env-default:"2"`

	DialogsURI string `env:"DIALOGS_URI"          env-default:"localhost:8080"      env-description:"IP:PORT GRPC URI"` //nolint:lll

	LogLevel string `env:"LOG_LEVEL" env-default:"info" env-description:"log level: trace, debug, info, warn, error, fatal, panic"` //nolint:lll
}

func issetEnvConfigFile() bool {
	_, err := os.Stat(envConfigFileName)

	return err == nil
}

func (e *EnvSetting) GetHelpString() (string, error) {
	customHeader := "options which can be set via env: "

	helpString, err := cleanenv.GetDescription(e, &customHeader)
	if err != nil {
		return "", fmt.Errorf("get help string failed: %w", err)
	}

	return helpString, nil
}

func New() *Config {
	envSetting := &EnvSetting{} //nolint:exhaustruct

	helpString, err := envSetting.GetHelpString()
	if err != nil {
		log.Panic("getting help string of env settings failed: ", err)
	}

	log.Info(helpString)

	if issetEnvConfigFile() {
		if err := cleanenv.ReadConfig(envConfigFileName, envSetting); err != nil {
			log.Panicf("read env cofig file failed: %s", err)
		}
	} else if err := cleanenv.ReadEnv(envSetting); err != nil {
		log.Panicf("read env config failed: %s", err)
	}

	return &Config{
		env: envSetting,
	}
}

func (c *Config) PrintDebug() {
	envReflect := reflect.Indirect(reflect.ValueOf(c.env))
	envReflectType := envReflect.Type()

	exp := regexp.MustCompile("([Tt]oken|[Pp]assword)")

	for i := 0; i < envReflect.NumField(); i++ {
		key := envReflectType.Field(i).Name

		if exp.MatchString(key) {
			val, _ := envReflect.Field(i).Interface().(string)
			log.Debugf("%s: len %d", key, len(val))

			continue
		}

		log.Debugf("%s: %v", key, spew.Sprintf("%#v", envReflect.Field(i).Interface()))
	}

	log.Infof("config loaded: %+v", *c.env)
}

func (c *Config) GetHost() string {
	return c.env.Host
}

func (c *Config) GetPort() string {
	return c.env.Port
}

func (c *Config) GetLogLevel() log.Level {
	lvl, err := log.ParseLevel(c.env.LogLevel)
	if err != nil {
		log.Error(err)

		return log.InfoLevel
	}

	return lvl
}

func (c *Config) GetClickhouseAddress() string {
	return fmt.Sprintf("%s:%d", c.env.ClickhouseHost, c.env.ClickhousePort)
}

func (c *Config) GetClickhouseUser() string {
	return c.env.ClickhouseUser
}

func (c *Config) GetShardsCount() int {
	return c.env.ShardCount
}

func (c *Config) GetDialogsURI() string {
	return c.env.DialogsURI
}

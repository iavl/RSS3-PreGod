package config

import (
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
)

type ProtocolStruct struct {
	Version string `koanf:"version"`
}

type HubServerStruct struct {
	RunMode      string        `koanf:"run_mode"`
	HttpPort     int           `koanf:"http_port"`
	ReadTimeout  time.Duration `koanf:"read_timeout"`
	WriteTimeout time.Duration `koanf:"write_timeout"`
}

type RedisStruct struct {
	Addr     string `koanf:"addr"`
	Password string `koanf:"password"`
	DB       int    `koanf:"db"`
}

type PostgresStruct struct {
	DSN             string        `koanf:"dsn"`
	MaxOpenConns    int           `koanf:"max_open_conns"`
	MaxIdleConns    int           `koanf:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `koanf:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `koanf:"conn_max_lifetime"`
}

type LoggerOutputConfig struct {
	Type     string `koanf:"type"`     // available values: `stdout`, `file`, `syslog`
	Filepath string `koanf:"filepath"` // only for file
	Facility int    `koanf:"facility"` // only for syslog, available values: `1 - 7` to `LOG_LOCAL0-LOG_LOCAL7`
}

type LoggerStruct struct {
	PrefixTag string `koanf:"prefix_tag"`
	Engine    string `koanf:"engine"`   // available values: `zap`
	Level     string `koanf:"level"`    // available values: `debug`, `info`, `warn`, `error`, `panic`, `fatal`
	Encoding  string `koanf:"encoding"` // available values: `json`, `console`

	Output []LoggerOutputConfig `koanf:"output"`
}

type ConfigStruct struct {
	Protocol  ProtocolStruct  `koanf:"protocol"`
	HubServer HubServerStruct `koanf:"hub_server"`
	Redis     RedisStruct     `koanf:"redis"`
	Postgres  PostgresStruct  `koanf:"postgres"`
	Logger    LoggerStruct    `koanf:"logger"`
}

var (
	Config = &ConfigStruct{}

	k = koanf.New(".")
)

func Setup() error {
	// Read user config
	if err := k.Load(file.Provider("config/config.dev.json"), json.Parser()); err != nil {
		return err
	}

	if err := k.Unmarshal("", Config); err != nil {
		return err
	}

	Config.HubServer.ReadTimeout = Config.HubServer.ReadTimeout * time.Second
	Config.HubServer.WriteTimeout = Config.HubServer.WriteTimeout * time.Second

	Config.Postgres.ConnMaxIdleTime = Config.Postgres.ConnMaxIdleTime * time.Second
	Config.Postgres.ConnMaxLifetime = Config.Postgres.ConnMaxLifetime * time.Second

	return nil
}

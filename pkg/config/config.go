package config

import (
	"fmt"
	"time"
)

// Config Object
// Stores environmant variable as configuration settings
type Config struct {
	Env      string `envconfig:"ENV" default:"development"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`
	LogDebug bool   `envconfig:"LOG_DEBUG" default:"false"`
	Server   struct {
		Host            string        `default:"0.0.0.0" envconfig:"HOST"`
		Port            string        `default:"8000" envconfig:"PORT"`
		DomainName      string        `default:"localhost" envconfig:"DOMAIN_NAME"`
		ReadTimeout     time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
		WriteTimeout    time.Duration `default:"20s" envconfig:"WRITE_TIMEOUT"`
		ShutdownTimeout time.Duration `default:"30s" envconfig:"SHUTDOWN_TIMEOUT"`
	}
	Mongo struct {
		URL         string        `envconfig:"MONGO_URL" default:"mongodb://localhost:27017"`
		Path        string        `envconfig:"MONGO_PATH" default:"database"`
		User        string        `envconfig:"MONGO_USER" default:""`
		Pass        string        `envconfig:"MONGO_PASS" default:""`
		DialTimeout time.Duration `envconfig:"DIAL_TIMEOUT" default:"30s"`
	}
	Redis struct {
		Host string `envconfig:"REDIS_HOST" default:"localhost"`
		Port string `envconfig:"REDIS_PORT" default:"6379"`
		Path string `envconfig:"REDIS_PATH" default:"0"`
	}
	RateLimit struct {
		Period time.Duration `default:"60m" envconfig:"RATE_DURATION"`
		Limit  int           `default:"1000" envconfig:"RATE_LIMIT"`
	}
	SMTP struct {
		Server   string `envconfig:"SMTP_SERVER" default:"smtp"`
		Port     int    `envconfig:"SMTP_PORT" default:"587"`
		User     string `envconfig:"SMTP_USER" default:"no-reply@cvcio.org"`
		From     string `envconfig:"SMTP_FROM" default:"no-reply@cvcio.org"`
		FromName string `envconfig:"SMTP_FROM_NAME" default:"MediaWatch"`
		Pass     string `envconfig:"SMTP_PASS" default:""`
		Reply    string `envconfig:"SMTP_REPLY" default:"info@cvcio.org"`
	}
}

// New Config create a new configuration object
func New() *Config {
	return new(Config)
}

// ServerURL returns server `host:port`
func (c *Config) ServerURL() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}

// RedisURL returns server `host:port`
func (c *Config) RedisURL() string {
	return fmt.Sprintf("%s:%s", c.Redis.Host, c.Redis.Port)
}

// RedisWithPathURL returns server `host:port`
func (c *Config) RedisWithPathURL() string {
	return fmt.Sprintf("redis://%s:%s/%s", c.Redis.Host, c.Redis.Port, c.Redis.Path)
}

// MongoURL returns server `host:port`
func (c *Config) MongoURL() string {
	return fmt.Sprintf("%s/%s", c.Mongo.URL, c.Mongo.Path)
}

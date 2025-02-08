package config

import (
	"reflect"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Debug         string        `mapstructure:"DEBUG"`
	ServerUrl     string        `mapstructure:"SERVER_URL"`
	DatabaseFile  string        `mapstructure:"DATABASE_FILE"`
	SessionConfig SessionConfig `mapstructure:"SESSION"`
}

type SessionConfig struct {
	CookieDomain string        `mapstructure:"COOKIE_DOMAIN"`
	CookieMaxAge time.Duration `mapstructure:"COOKIE_MAX_AGE"`
	CookieName   string        `mapstructure:"COOKIE_NAME"`
	CookieSecure bool          `mapstructure:"COOKIE_SECURE"`
	Key          string        `mapstructure:"KEY"`
}

func BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			BindEnvs(v.Interface(), append(parts, tv)...)
		default:
			_ = viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

func LoadConfig() (Config, error) {
	var cfg Config

	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	BindEnvs(cfg)

	err := viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

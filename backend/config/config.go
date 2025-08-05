package config

import (
	"strings"

	"github.com/spf13/viper"
	"github.com/torbenconto/spooler/internal/types"
)

var Cfg *Config

type Config struct {
	Port             int      `mapstructure:"port"`
	SecretKey        string   `mapstructure:"secret_key"`
	Mode             string   `mapstructure:"mode"`
	CORSAllowOrigins []string `mapstructure:"cors_allow_origins"`

	Features struct {
		EmailWhitelistEnabled bool `mapstructure:"email_whitelist_enabled"`
	} `mapstructure:"features"`

	Supabase struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Database string `mapstructure:"database"`
	} `mapstructure:"supabase"`

	SMTP struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Email    string `mapstructure:"email"`
		Password string `mapstructure:"password"`
	} `mapstructure:"smtp"`

	Admin struct {
		Email     string `mapstructure:"email"`
		FirstName string `mapstructure:"first_name"`
		LastName  string `mapstructure:"last_name"`
	} `mapstructure:"admin"`

	Storage struct {
		Provider types.StorageProvider `mapstructure:"provider"`

		GoogleCloud struct {
			BucketName string `mapstructure:"bucket_name"`
		} `mapstructure:"google_cloud"`

		Local struct {
			BasePath string `mapstructure:"base_path"`
		} `mapstructure:"local"`
	} `mapstructure:"storage"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	_ = viper.BindEnv("port")
	_ = viper.BindEnv("secret_key")
	_ = viper.BindEnv("mode")
	_ = viper.BindEnv("email_whitelist_enabled")
	_ = viper.BindEnv("cors_allow_origins")

	_ = viper.BindEnv("features.email_whitelist_enabled")

	_ = viper.BindEnv("supabase.host")
	_ = viper.BindEnv("supabase.port")
	_ = viper.BindEnv("supabase.user")
	_ = viper.BindEnv("supabase.password")
	_ = viper.BindEnv("supabase.database")

	_ = viper.BindEnv("smtp.host")
	_ = viper.BindEnv("smtp.port")
	_ = viper.BindEnv("smtp.email")
	_ = viper.BindEnv("smtp.password")

	_ = viper.BindEnv("admin.email")
	_ = viper.BindEnv("admin.first_name")
	_ = viper.BindEnv("admin.last_name")

	_ = viper.BindEnv("storage.provider")
	_ = viper.BindEnv("storage.google_cloud.bucket_name")
	_ = viper.BindEnv("storage.local.base_path")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	Cfg = &cfg
	return &cfg, nil
}

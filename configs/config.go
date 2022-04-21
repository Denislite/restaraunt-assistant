package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"time"
)

type (
	Config struct {
		Environment string
		Postgres    PostgresConfig
		FileStorage FileStorageConfig
		HTTP        HTTPConfig
		GRPC        GRPCConfig
		GRPCFD      GRPCConfigFD
		GRPCCS      GRPCConfigCS
		GRPCAuth    GRPCConfigAuth
	}

	PostgresConfig struct {
		Port     string
		Sslmode  string
		Host     string
		Username string
		Dbname   string
		Password string
	}

	FileStorageConfig struct {
		Endpoint  string
		Bucket    string
		AccessKey string
		SecretKey string
	}

	HTTPConfig struct {
		Host               string
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}

	GRPCConfig struct {
		Port string
	}

	GRPCConfigFD struct {
		Host string
		Port string
	}

	GRPCConfigAuth struct {
		Host string
		Port string
	}

	GRPCConfigCS struct {
		Host string
		Port string
	}
)

func Init(configsDir string) (*Config, error) {

	if err := parseConfigFile(configsDir); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {

	if err := viper.UnmarshalKey("db", &cfg.Postgres); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("fileStorage", &cfg.FileStorage); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpcFD", &cfg.GRPCFD); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpcCS", &cfg.GRPCCS); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("grpcAuth", &cfg.GRPCAuth); err != nil {
		return err
	}
	return nil
}

func setFromEnv(cfg *Config) {
	_ = godotenv.Load()
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")

	cfg.FileStorage.AccessKey = os.Getenv("STORAGE_ACCESS_KEY")
	cfg.FileStorage.SecretKey = os.Getenv("STORAGE_SECRET_KEY")
}

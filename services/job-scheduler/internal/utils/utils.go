package utils

import (
	"fmt"
	"job-scheduler/internal/config"
)

func CreatePostgresConnectionString(cfg *config.Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBConfig.Credentials.Username, cfg.DBConfig.Credentials.Password, cfg.DBConfig.Credentials.Host, cfg.DBConfig.Credentials.Port, cfg.DBConfig.Credentials.Name, cfg.DBConfig.Credentials.SSLMode)

}

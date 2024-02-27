package main

import (
	"encoding/json"
	"job-scheduler/internal/config"
	"job-scheduler/pkg/logging"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap/zapcore"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:  "",
	Long: "Root cmd",
	Run: func(cmd *cobra.Command, args []string) {
		runApplication()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "conf", "", "", "config file path")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Printf("failed to execute command. err: %v", err)
		os.Exit(1)
	}
}

func runApplication() {
	// load configs and sets default logger configs.
	conf, err := config.Load(configFile)
	if err != nil {
		log.Fatal(err)
	}
	logging.SetConfig(&logging.Config{
		Encoding:    conf.LoggingConfig.Encoding,
		Level:       zapcore.Level(conf.LoggingConfig.Level),
		Development: conf.LoggingConfig.Development,
	})
	defer logging.DefaultLogger().Sync()

	// setup application(di + run server)

}

func printAppInfo(cfg *config.Config) {
	b, _ := json.MarshalIndent(&cfg, "", "  ")
	logging.DefaultLogger().Infof("application information\n%s", string(b))
}

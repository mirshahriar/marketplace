package cmd

import (
	"log"

	"github.com/mirshahriar/marketplace/config"
	app "github.com/mirshahriar/marketplace/internal/application"
	"github.com/mirshahriar/marketplace/internal/framework/http"
	db "github.com/mirshahriar/marketplace/internal/framework/mysql"
	"github.com/spf13/cobra"
)

// Run ...
// @title Marketplace
// @version 1.0
// @description API Documentation for Marketplace
// @contact.name API Support
// @host localhost:8080
// @BasePath /
// @schemes http
var runCmd = &cobra.Command{
	Use:   "serve",
	Short: "run marketplace server",
	Run: func(cmd *cobra.Command, args []string) {
		dbConfig := config.GetDBConfig()
		dbAdapter, err := db.NewAdapterWithConfig(dbConfig)
		if err != nil {
			log.Fatal(err)
		}

		appConfig := config.GetAppConfig()

		application := app.NewApplication(
			appConfig,
			dbAdapter,
		)

		// Creating a new http adapter with the application.
		httpAdapter := http.NewAdapter(appConfig, application)
		// Running the http server.
		httpAdapter.Run()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Initializing the config.
		config.Init()
	},
}

func init() {
	// Adding the run command to the root command.
	rootCmd.AddCommand(runCmd)
}

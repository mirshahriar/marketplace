package cmd

import (
	"log"

	"github.com/mirshahriar/marketplace/config"
	db "github.com/mirshahriar/marketplace/internal/framework/mysql"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate mysql database",
	Run: func(cmd *cobra.Command, args []string) {
		dbConfig := config.GetDBConfig()
		dbAdapter, err := db.NewAdapterWithConfig(dbConfig)
		if err != nil {
			log.Fatal(err.Error())
		}

		if cErr := dbAdapter.Migrate(); cErr != nil {
			log.Println(cErr.Error())
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		// Initializing the config.
		config.Init()
	},
}

func init() {
	// Adding the run command to the root command.
	rootCmd.AddCommand(migrateCmd)
}

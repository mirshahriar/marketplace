package cmd

import (
	"log"
	"os"

	"github.com/mirshahriar/marketplace/config"
	db "github.com/mirshahriar/marketplace/internal/framework/mysql"
	"github.com/spf13/cobra"
)

func getMigrateCmd() *cobra.Command {
	var generateFakeUser bool
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "migrate mysql database",
		Run: func(cmd *cobra.Command, args []string) {
			appConfig := config.GetAppConfig()

			dbConfig := config.GetDBConfig()
			dbAdapter, err := db.NewAdapterWithConfig(appConfig, dbConfig)
			if err != nil {
				log.Fatal(err.Error())
			}

			if cErr := dbAdapter.Migrate(); cErr != nil {
				log.Println(cErr.Error())
			}

			if generateFakeUser {
				fakeUser, err2 := dbAdapter.CreateFakeUser()
				if err2 != nil {
					log.Println(err2.Error())
					os.Exit(1)
				}

				log.Println("Fake user created successfully")
				log.Println("ID: ", fakeUser.UserID)
				log.Println("Token: ", fakeUser.Token)
			}
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			// Initializing the config.
			config.Init()
		},
	}

	migrateCmd.Flags().BoolVarP(&generateFakeUser, "add-fake-user", "f", false, "fake user data")

	return migrateCmd
}

func init() {
	// Adding the run command to the root command.
	rootCmd.AddCommand(getMigrateCmd())
}

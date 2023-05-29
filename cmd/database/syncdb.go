package database

import (
	"RudderMaster/database"
	"RudderMaster/settings"
	"RudderMaster/utils/orm"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	confFile string

	SyncDBCmd = &cobra.Command{
		Use:   "syncdb",
		Short: "sync database table",
		Long:  "use config to migrate db tables",
		PreRun: func(cmd *cobra.Command, args []string) {
			setDatabase()
		},
		Run: func(cmd *cobra.Command, args []string) {
			migrateDatabase()
		},
	}
)

func init() {
	SyncDBCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "./config/config.yml", "config file default ./config/config.yml")
}

func setDatabase() {
	settings.SetupConfig(confFile)
	database.SetupDB()
}

func migrateDatabase() {
	defer func() {
		sqlDb, _ := database.DB.DB()
		err := sqlDb.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	orm.DataBaseAutoMigrates(orm.TableModes...)
}

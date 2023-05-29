package database

import (
	"RudderMaster/database"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	configFile    string
	username      string
	password      string
	isAdmin       bool
	CreateUserCmd = &cobra.Command{
		Use:   "createuser",
		Short: "create user",
		Long:  "use cmd createuser to create a user. -u username -p password -a is_admin",
		PreRun: func(cmd *cobra.Command, args []string) {
			setDatabase()
			// TODO: 验证参数
		},
		Run: func(cmd *cobra.Command, args []string) {
			createUser()
		},
	}
)

func init() {
	CreateUserCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./config/config.yml", "config file default ./config/config.yml")
	CreateUserCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "please input username")
	CreateUserCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "please input password")
	CreateUserCmd.PersistentFlags().BoolVarP(&isAdmin, "is_admin", "a", false, "use -a [true|false] to set is admin")
}

func createUser() {
	defer func() {
		sqlDb, _ := database.DB.DB()
		err := sqlDb.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("to Create User!")
}

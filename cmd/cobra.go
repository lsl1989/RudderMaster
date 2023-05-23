package cmd

import (
	"RudderMaster/cmd/server"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "RudderMaster",
	Short: "rudder master",
	Long:  "rudder master",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least on arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		useStr := "欢迎使用 RudderMaster，可以使用 -h 查看帮助"
		fmt.Println(useStr)
	},
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

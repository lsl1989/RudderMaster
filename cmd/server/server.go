package server

import (
	"RudderMaster/database"
	"RudderMaster/router"
	"RudderMaster/settings"
	"RudderMaster/utils/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var (
	mode     string
	confFile string

	StartCmd = &cobra.Command{
		Use:   "server",
		Short: "start server",
		Long:  "start server user args",
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("正在启动server")
			// TODO: 加载配置 初始化数据库等
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "c", "./config/config.yml", "config file default ./config/config.yml")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "m", gin.DebugMode, "server mode ; eg:dev,test,prod")
}

func setup() {
	// 加载配置文件
	settings.SetupConfig(confFile)
	// 初始化数据库
	database.SetupDB()
}

func run() error {
	handle := router.InitRouter()
	defer func() {
		sqlDb, _ := database.DB.DB()
		err := sqlDb.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	appConfig := settings.Config.Application
	timeout := time.Second * time.Duration(appConfig.TimeOut)
	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", appConfig.Host, appConfig.Port),
		Handler:      handle,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}
	fmt.Println(fmt.Sprintf("Running http://%s:%s", appConfig.Host, appConfig.Port))
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

	//fmt.Println(settings.Config)
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	return nil
}

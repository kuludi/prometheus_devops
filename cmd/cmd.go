package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"prometheus_devops/config"
	"prometheus_devops/db"
	"prometheus_devops/model"
	"prometheus_devops/routers"
)

var (
	cfgFile string
	mode    string
	rootCmd = &cobra.Command{}
)

func initConfig() {
	config.MustInit(os.Stdout, cfgFile)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config/config.yaml", "config file")
	rootCmd.PersistentFlags().StringVar(&mode, "mode", "debug", "Set gin mode")
	//viper.SetDefault("gin.mode", rootCmd.PersistentFlags().Lookup("debug"))
}

func Execute() error {
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		_, err := db.MySql(
			viper.GetString("db.hostname"),
			viper.GetString("db.port"),
			viper.GetString("db.username"),
			viper.GetString("db.password"),
			viper.GetString("db.dbname"),
		)
		if err != nil {
			return err
		}

		db.DB.AutoMigrate(&model.User{})



		r := routers.SerRouter(mode)

		return r.Run(viper.GetString("server.port"))

	}

	return rootCmd.Execute()
}

/*
Copyright © 2024 NAME HERE rnikrozoft.dev@gmail.com
*/
package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/rnikrozoft/himymonsters-master-backend/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "himymonsters-master-backend",
	Short: "This is a backend service for himymonsters game master",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var conf *config.Config

func init() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("cannot unmarshal: %w", err))
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.App.Database)))
	config.Database = bun.NewDB(sqldb, pgdialect.New())
}

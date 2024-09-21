/*
Copyright © 2024 NAME HERE rnikrozoft.dev@gmail.com
*/
package cmd

import (
	"github.com/rnikrozoft/himymonsters-master-backend/config"
	"github.com/rnikrozoft/himymonsters-master-backend/repository"
	"github.com/rnikrozoft/himymonsters-master-backend/service"
	"github.com/spf13/cobra"
)

// shopCmd represents the shop command
var shopCmd = &cobra.Command{
	Use:   "shop",
	Short: "Shop menagements",
	Long:  `this command is used to manage the shop items. like adding new items to the shop, updating the existing items, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		if a, _ := cmd.Flags().GetBool("add"); a {
			repo := repository.NewRepository(cmd.Context(), config.Database)
			if err := service.NewShop(repo, conf.Shop.ID, conf.Shop.Sheet, conf.Shop.Range).AddItems(); err != nil {
				cmd.PrintErrln(err)
			}
		} else {
			cmd.Help()
		}

	},
}

func init() {
	rootCmd.AddCommand(shopCmd)
	shopCmd.Flags().BoolP("add", "a", false, "Add items to the shop")
}

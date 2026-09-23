/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"adhoc/internal/services"
	"adhoc/internal/store"
	"database/sql"
	"os"

	"github.com/spf13/cobra"
)

var (
	db          *sql.DB
	IdeaService *services.IdeaService
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "adhoc",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		db, err := store.InitDatabase()
		if err != nil {
			return err
		}
		ideaStore := store.New(db)
		IdeaService = services.NewIdeaService(ideaStore)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

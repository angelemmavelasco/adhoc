package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	titleFlag    string
	contentFlag  string
	keywordsFlag []string
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Record a new idea on the fly",
	Long:  `Create and ingest a new thought into the database, automatically linking keywords.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		created, err := IdeaService.IngestIdea(titleFlag, contentFlag, keywordsFlag)
		if err != nil {
			return fmt.Errorf("error while saving idea: %w", err)
		}

		color.Green("✓ Idea saved correctly: '%s' (ID: %d)\n", created.Title, created.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&titleFlag, "title", "t", "", "Title of the idea")
	newCmd.Flags().StringVarP(&contentFlag, "content", "c", "", "Content or description")
	newCmd.Flags().StringSliceVarP(&keywordsFlag, "keywords", "k", []string{}, "Keywords separated by commas")

	_ = newCmd.MarkFlagRequired("title")
	_ = newCmd.MarkFlagRequired("content")
}

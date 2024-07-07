package cmd

import (
	"github.com/akifkadioglu/organizer/config"
	"github.com/akifkadioglu/organizer/olog"
	"github.com/akifkadioglu/organizer/run"
	"github.com/spf13/cobra"
)

// categorizeCmd represents the categorize command
var categorizeCmd = &cobra.Command{
	Use:     "categorize",
	Version: config.ReadValue().AppVersion,
	Example: "organizer categorize --path=/path/to/folder",
	Short:   "organizer categorize --path=/path/to/folder | default path is current directory",
	Long:    "Organizer looks paths files type and folders your files.",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		if path == "" {
			path = "."
		}
		res, err := run.Categorize(path)
		if err != nil {
			olog.Error(err.Error())
		} else {
			olog.Success(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(categorizeCmd)
	categorizeCmd.PersistentFlags().String("path", "", "The path to categorize files. Default is current directory.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categorizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categorizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

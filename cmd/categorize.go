package cmd

import (
	"strings"

	"github.com/akifkadioglu/organizer/config"
	"github.com/akifkadioglu/organizer/olog"
	"github.com/akifkadioglu/organizer/run"
	"github.com/spf13/cobra"
)

// categorizeCmd represents the categorize command
var categorizeCmd = &cobra.Command{
	Use:     "categorize",
	Version: config.ReadValue().AppVersion,
	Example: "organizer categorize /path/to/folder",
	Short:   "organizer categorize /path/to/folder | default path is current directory",
	Long:    "Document Organizer looks paths files type and folders your files.",
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) != 0 {
			path = ""
			for _, v := range args {
				path += " " + v
			}
		}
		path = strings.TrimSpace(path)
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categorizeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categorizeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

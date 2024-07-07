/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/akifkadioglu/organizer/olog"
	"github.com/akifkadioglu/organizer/run"
	"github.com/spf13/cobra"
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "organizer password password_name",
	Long:  "Organizer password is a password manager. You can generate and save your passwords with this command.",
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetString("username")
		app, _ := cmd.Flags().GetString("app")
		if len(args) != 0 {
			arg := args[0]
			switch arg {
			case "remove":
				run.RemovePassword(app, name)
			case "list":
				run.ListPassword(app)
			case "generate":
				run.GeneratePassword(app, name)
			case "get":
				run.GetPassword(app, name)
			case "json":
				run.PasswordDownload()
			default:
				olog.Error("Invalid argument")
			}
		} else {
			olog.Error("Please enter a valid argument.")
			olog.Fatal("You can use generate | get | list | remove | json")
		}
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	passwordCmd.PersistentFlags().String("username", "", "A string to save password with a username")
	passwordCmd.PersistentFlags().String("app", "", "A string to save password to an app name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// passwordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// passwordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

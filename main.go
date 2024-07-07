package main

import (
	"github.com/akifkadioglu/organizer/cmd"
	"github.com/akifkadioglu/organizer/database"
)

func main() {
	database.Connect()
	cmd.Execute()
}

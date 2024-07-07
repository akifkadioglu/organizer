package run

import (
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/akifkadioglu/organizer/config"
	"github.com/akifkadioglu/organizer/olog"
	"github.com/olekukonko/tablewriter"
)

// checkPath checks if the path exists
func checkPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		olog.Fatal("Path does not exist")
	}
}

// createFolders creates folders for each type
func createFolders(path string) error {
	var err error

	folders := []string{
		config.ReadValue().AudioFolder,
		config.ReadValue().VideoFolder,
		config.ReadValue().ImageFolder,
		config.ReadValue().ApplicationFolder,
		config.ReadValue().OfficeFolder,
		config.ReadValue().OthersFolder,
	}

	for _, folder := range folders {
		err := os.Mkdir(filepath.Join(path, folder), 0755)
		if err != nil && os.IsNotExist(err) {
			return err
		}
	}

	return err
}

// moveDocs moves documents to their respective folders
func moveDocs(all []fs.DirEntry, path string) {
	for _, element := range all {

		if element.IsDir() {
			continue
		}

		extenstion := strings.Split(element.Name(), ".")
		destPath := ""

		if len(extenstion) == 1 {
			destPath = path + "/" + config.ReadValue().OthersFolder
		} else {
			destPath = path + "/" + getFolder(extenstion[1])
		}

		os.Rename(path+"/"+element.Name(), destPath+"/"+element.Name())

	}
}

// getFolder returns the folder name for the given extension
func getFolder(extenstion string) string {

	switch {
	case slices.Contains(config.ReadValue().AudioTypes, extenstion):
		return config.ReadValue().AudioFolder
	case slices.Contains(config.ReadValue().VideoTypes, extenstion):
		return config.ReadValue().VideoFolder
	case slices.Contains(config.ReadValue().ImageTypes, extenstion):
		return config.ReadValue().ImageFolder
	case slices.Contains(config.ReadValue().ApplicationTypes, extenstion):
		return config.ReadValue().ApplicationFolder
	case slices.Contains(config.ReadValue().OfficeTypes, extenstion):
		return config.ReadValue().OfficeFolder
	default:
		return config.ReadValue().OthersFolder
	}

}

// informationText creates a text file with information about the types
func informationText(path string) error {
	file, err := os.Create(path + "/types.txt")

	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(
		"Audio Types: " + strings.Join(config.ReadValue().AudioTypes, ", ") +
			"\n" + "Video Types: " + strings.Join(config.ReadValue().VideoTypes, ", ") +
			"\n" + "Image Types: " + strings.Join(config.ReadValue().ImageTypes, ", ") +
			"\n" + "Application Types: " + strings.Join(config.ReadValue().ApplicationTypes, ", ") +
			"\n" + "Office Types: " + strings.Join(config.ReadValue().OfficeTypes, ", ") +
			"\n" + "Others Types: Literally everything else")
	return nil
}

// Log Password as table
func logPasswordTable(header []string, data [][]string) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.AppendBulk(data)
	table.Render()
}
package run

import (
	"flag"
	"os"
)

// Categorize is a function that categorizes files in a given path
func Categorize(path string) (string, error) {
	var err error

	checkPath(path)

	err = createFolders(path)
	if err != nil {
		return "", err
	}

	pathPtr := flag.String("path", path, "")
	flag.Parse()

	if _, err := os.Stat(*pathPtr); os.IsNotExist(err) {
		return "", err
	}

	all, err := os.ReadDir(*pathPtr)
	if err != nil {
		return "", err
	}

	moveDocs(all, path)

	err = informationText(path)
	if err != nil {
		return "", err
	}

	return "organized successfully", nil
}

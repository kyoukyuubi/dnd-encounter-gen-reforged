package errorhandling

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// the structure of the logs file
type ErrorFile struct {
	Time             time.Time `json:"time"`
	Err              string    `json:"error"`
	TriggeredCommand string    `json:"triggeredCommand"`
}

func folderCheck() {
	// set the parent folder
	parentFolder := "json"

	// use filepath to make the filepath
	logsFolder := filepath.Join(parentFolder, "logs")

	// check if folder exists, if it doesn't create it
	if err := os.MkdirAll(logsFolder, 0755); err != nil {
		panic(err)
	}
}

func LogError(inputErr error, cmd string) {
	// run the folderCheck() function to check if the folder exists
	folderCheck()

	// make the struct for the file
	fileData := ErrorFile{
		Time:             time.Now(),
		Err:              inputErr.Error(),
		TriggeredCommand: cmd,
	}

	// set the filename
	fileName := time.Now().String() + "_error.json"

	// set the fullpath
	fullpath := "json/logs/" + fileName

	// create the file
	file, err := os.Create(fullpath)
	if err != nil {
		panic(err)
	}

	// make sure it closes after function returns
	defer file.Close()

	// write to the file
	if err := json.NewEncoder(file).Encode(fileData); err != nil {
		panic(err)
	}
}

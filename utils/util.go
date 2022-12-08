package utils

import (
	"io"
	"os"
	"fmt"
	"time"
	"strings"
	"runtime"

	gonLogger "github.com/mellowarex/gonc/logger"
)
func UnknownSubCommand(subCommand string) {
	errMsg := `
gonc: "` + subCommand +`" Unknown subcommand.
Use gonc help for more information.`
	io.WriteString(os.Stdout, errMsg)
	os.Exit(2)
}

func PrintErrorAndExit(errorMsg string) {
	errMsg :=`
gonc: ` + errorMsg + `
Use gonc help for more information.`
	io.WriteString(os.Stdout, errMsg)
	os.Exit(2)
}

func GetWorkPath() string {
	currentPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return currentPath
}

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// askForConfirmation uses Scanln to parse user input. A user must type in "yes" or "no" and
// then press enter. It has fuzzy matching, so "y", "Y", "yes", "YES", and "Yes" all count as
// confirmations. If the input is not recognized, it will ask again. The function does not return
// until it gets a valid response from the user. Typically, you should use fmt to print out a question
// before calling askForConfirmation. E.g. fmt.Println("WARNING: Are you sure? (yes/no)")
func AskForConfirmation() bool {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		gonLogger.Log.Fatalf("%s", err)
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		return AskForConfirmation()
	}
}

func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		if elem == element {
			return true
		}
	}
	return false
}

// WriteToFile creates a file and writes content to it
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func GetGoVersionSkipMinor() string {
	strArray := strings.Split(runtime.Version()[2:], `.`)
	return strArray[0] + `.` + strArray[1]
}

// __FILE__ returns the file name in which the function was invoked
func FILE() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// __LINE__ returns the line number at which the function was invoked
func LINE() int {
	_, _, line, _ := runtime.Caller(1)
	return line
}

// GetFileModTime returns unix timestamp of `os.File.ModTime` for the given path.
func GetFileModTime(path string) int64 {
	path = strings.Replace(path, "\\", "/", -1)
	f, err := os.Open(path)
	if err != nil {
		gonLogger.Log.Errorf("Failed to open file on '%s': %s", path, err)
		return time.Now().Unix()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		gonLogger.Log.Errorf("Failed to get file stats: %s", err)
		return time.Now().Unix()
	}

	return fi.ModTime().Unix()
}
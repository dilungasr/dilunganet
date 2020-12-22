package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

var colorGreen = "\033[32m"
var colorReset = "\033[0m"

func main() {
	path := os.Getenv("Path")
	workingDirectory, _ := os.Getwd()

	// check if doen installed
	allPathValues := strings.Split(path, ";")
	alreadyInstalled := false

	for _, v := range allPathValues {
		if v == workingDirectory {
			alreadyInstalled = true
		}
	}

	if alreadyInstalled {
		fmt.Println("Dilunga Net was already added to you computer....")
		time.Sleep(3 * time.Second)
		fmt.Println("YOU CAN simply run Dilunga Net by using dnet command")

	} else {
		setPath := exec.Command("setx", "Path", path+";"+workingDirectory)

		if err := CommandComplete(setPath); err != nil {
			fmt.Println("Error installing Dilunga Net: " + err.Error())
		}

		fmt.Printf("Succefully installed Dilunga Net...")
		time.Sleep(3 * time.Second)
		fmt.Println("YOU CAN NOW run Dilunga Net by using dnet command")
	}
	time.Sleep(30 * time.Second)

}

// CommandComplete for completing command
func CommandComplete(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	// run the command
	return cmd.Run()
}

package actions

import (
	"fmt"
	"os"
	"os/exec"
)

// The color for terminals which supports colored texts
var colorGreen = "\033[32m"
var colorCyan = "\033[36m"
var colorYellow = "\033[33m"
var colorReset = "\033[0m"

// Start is a function for starting the hostednetwork
func Start() error {
	startHostedNetwork := exec.Command("netsh", "wlan", "start", "hostednetwork")

	err := CommandComplete(startHostedNetwork)

	if err != nil {
		return err
	}

	// clear all the output
	if err := Clear(); err != nil {
		return err
	}

	// print success output to the console
	fmt.Print(string(colorGreen), `
			  CONGRATS!!! Your network is ready! :)
________________________________________________________________________
			            By `,
	)
	fmt.Println(string(colorCyan), "Dilunga", string(colorReset))

	return nil
}

// CommandComplete for completing command
func CommandComplete(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	// run the command
	return cmd.Run()
}

// Clear for clearing the console
func Clear() error {
	clearInCmdAndPowershellBased := exec.Command("cmd", "/c", "cls")
	err := CommandComplete(clearInCmdAndPowershellBased)
	if err != nil {
		fmt.Println(err)
		// run bash based commands
		clearInBashBased := exec.Command("clear")

		err = CommandComplete(clearInBashBased)
		if err != nil {
			return err
		}
	}

	return nil
}

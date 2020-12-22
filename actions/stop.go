package actions

import (
	"fmt"
	"os"
	"os/exec"
)

// Stop is a function to stop the hostednetwork
func Stop() error {
	stopHostedNetwork := exec.Command("netsh", "wlan", "stop", "hostednetwork")

	stopHostedNetwork.Stdout = os.Stdout
	stopHostedNetwork.Stderr = os.Stdout

	//run the command
	if err := stopHostedNetwork.Run(); err != nil {
		return err
	}

	// Clear the console.
	Clear()

	// Goodbye message....
	fmt.Println(string(colorGreen), `
	                  See you again ;)`, string(colorReset),
	)

	return nil
}

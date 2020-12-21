package actions

import (
	"fmt"
	"os"
	"os/exec"
)

// Stop is a function to stop the hostednetwork
func Stop() {
	stopHostedNetwork := exec.Command("netsh", "wlan", "stop", "hostednetwork")

	stopHostedNetwork.Stdout = os.Stdout
	stopHostedNetwork.Stderr = os.Stdout

	//run the command
	if err := stopHostedNetwork.Run(); err != nil {
		fmt.Println(err)
	}
}

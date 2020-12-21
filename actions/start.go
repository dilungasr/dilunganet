package actions

import (
	"fmt"
	"os"
	"os/exec"
)

// Start is a function for starting the hostednetwork
func Start() {
	// The green color for terminals which supports colored texts
	colorGreen := "\033[32m"
	startHostedNetwork := exec.Command("netsh", "wlan", "start", "hostednetwork")

	startHostedNetwork.Stdout = os.Stdout
	startHostedNetwork.Stderr = os.Stdout

	if err := startHostedNetwork.Run(); err != nil {
		fmt.Println(err)
		return
	}

	// print success output to the console
	fmt.Println(string(colorGreen), `
			  CONGRATS!!! Your network is ready! :)
________________________________________________________________________
			            By Dilunga`,
	)

}

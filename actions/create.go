package actions

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Create is function for creating the hostednetworks
func Create() {
	// take the name of the network
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your network name: ")
	netname, err := reader.ReadString('\r')
	netname = strings.Replace(netname, "\n", "", -1)
	netname = strings.Replace(netname, "\r", "", -1)

	if err != nil {
		fmt.Println(err)
		return
	}

	// take the password of the network
	fmt.Print("Enter your network password: ")
	reader2 := bufio.NewReader(os.Stdin)
	netpassword, err := reader2.ReadString('\r')
	netpassword = strings.Replace(netpassword, "\n", "", -1)
	netpassword = strings.Replace(netpassword, "\r", "", -1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// let's creafte the network
	createNetwork := exec.Command("netsh", "wlan", "set", "hostednetwork", "mode=allow", "ssid="+netname, "key="+netpassword)

	// set where to ouput errors and results
	createNetwork.Stdout = os.Stdout
	createNetwork.Stderr = os.Stderr

	// execute the command
	if err = createNetwork.Run(); err != nil {
		fmt.Println(err)
		return
	}
}

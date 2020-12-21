package main

import (
	"fmt"
	"strings"

	"github.com/windroldev/dilunganet/actions"
)

func main() {
	var action string
	actionController := 0

	for actionController == 0 {
		fmt.Print("What do you want to do? (stop/start/create/install): ")
		action = ""
		if _, err := fmt.Scanf("%s", &action); err != nil {
			fmt.Println(err)
			return
		}
		if strings.ToLower(action) == "stop" || strings.ToLower(action) == "start" || strings.ToLower(action) == "create" {
			actionController = 1
		} else {
			fmt.Println("Please choose the proper action")
		}

	}

	// if it requires to stop the started hostednetwork
	if strings.ToLower(action) == "stop" {
		actions.Stop()
		// end the program
		return
	}
	// if it requires to only start the hostednetwork without recreating a new one
	if strings.ToLower(action) == "start" {
		actions.Start()
		// end the program
		return
	}

	// if it requires to create a new virtual wlan hostednetwork
	if strings.ToLower(action) == "create" {
		actions.Create()
		actions.Start()
	}

}

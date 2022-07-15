package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func logger() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("Press ESC to quit")
	currentStr := ""
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}

		switch key {
		case keyboard.KeySpace:
			file.WriteString(currentStr + "\n")
			currentStr = ""
			continue
		case keyboard.KeyEnter:
			file.WriteString(currentStr + "\n")
			currentStr = ""
			continue
		case keyboard.KeyTab:
			file.WriteString(currentStr + "\n")
			currentStr = ""
			continue
		}
		formattedStr := fmt.Sprintf("%q", char)[1 : len(fmt.Sprintf("%q", char))-1]
		currentStr = currentStr + formattedStr
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
	}
}

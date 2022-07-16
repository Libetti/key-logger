package keyboardlogger

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func StartKeyboardLogger() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0644)
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
			file.WriteString(fmt.Sprintf("%s: WriteChar = %s: String = %s \n \n", time.Now().String(), "key_Space", currentStr))
			currentStr = ""
			continue
		case keyboard.KeyEnter:
			file.WriteString(fmt.Sprintf("%s: WriteChar = %s: String = %s \n \n", time.Now().String(), "key_Enter", currentStr))
			currentStr = ""
			continue
		case keyboard.KeyTab:
			file.WriteString(fmt.Sprintf("%s: WriteChar = %s: String = %s \n \n", time.Now().String(), "key_Tab", currentStr))
			currentStr = ""
			continue
		}
		formattedStr := fmt.Sprintf("%q", char)[1 : len(fmt.Sprintf("%q", char))-1]
		currentStr = currentStr + formattedStr
		fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
	}
}

func StopKeyboardLogger() {
	_ = keyboard.Close()
}

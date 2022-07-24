package keyboardlogger

import (
	"fmt"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func StartKeyboardLogger(logsDirectory string, file *os.File, loggerRunning *bool) {
	if err := keyboard.Open(); err != nil {
		*loggerRunning = false
		return
	}
	*loggerRunning = true
	currentStr := ""
	fmt.Printf("On \n")

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			StopKeyboardLogger(loggerRunning)
			*loggerRunning = false
			if len(currentStr) > 0 {
				file.WriteString(fmt.Sprintf("%s: WriteChar = %s: String = %s \n \n", time.Now().String(), "keyboard_terminated", currentStr))

			}
			break
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
	}
}

func StopKeyboardLogger(loggerRunning *bool) {
	fmt.Printf("Off \n")

	_ = keyboard.Close()
	*loggerRunning = false
}

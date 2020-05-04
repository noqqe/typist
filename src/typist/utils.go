package typist

import "fmt"

// Colored output on commandline
func LogMessage(message string, color string) {
	if color == "red" {
		fmt.Printf("%s%s%s\n", colorRed, message, colorReset)
	} else if color == "green" {
		fmt.Printf("%s%s%s\n", colorGreen, message, colorReset)
	} else {
		fmt.Printf("%s\n", message)
	}
}

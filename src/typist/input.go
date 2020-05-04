package typist

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

type typeResponse struct {
	failure bool
	quit    bool
	char    string
}

// Read input from Keystrokes and compare with expected input
func readInput(expect string) typeResponse {

	var resp typeResponse
	resp.failure = false
	resp.quit = false

	char, key, err := keyboard.GetKey()
	if err != nil {
		LogMessage(fmt.Sprintf("GetKey: %v ", err), "red")
		os.Exit(1)
	}

	// check if non alpha numeric input (like space)
	// fmt.Print("\n", int(key), int([]rune(expect)[0]), "\n")
	if char == 0 {

		// check user wants to quit
		if int(key) == 3 {
			resp.quit = true
			resp.char = "abort...\n"
			return resp
		}

		// compare if real key like space
		if int(key) == int([]rune(expect)[0]) {
			resp.failure = false
			resp.char = string(key)
			return resp
		}
	}

	// fmt.Println(string(char), expect)
	if string(char) == expect {
		resp.char = string(char)
		return resp
	} else {
		resp.failure = true
		resp.char = fmt.Sprintf("%sx%s", colorRed, colorReset)
		return resp
	}

}

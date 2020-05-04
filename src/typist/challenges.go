package typist

import (
	"fmt"
	"time"
)

const (
	challengeLimiter string = ">> "
	colorCyan        string = "\033[36m"
	colorGreen       string = "\033[32m"
	colorPurple      string = "\033[35m"
	colorRed         string = "\033[31m"
	colorReset       string = "\033[0m"
	failureChar             = 10008 // cross heavy utf8
	successChar             = 10004 // checkmark utf8
	version          string = "0.0.2"
)

// Prompt user the challenge and return if succeeded and what percentage
func ShowChallenge(challenge string) (int, float64) {

	var elapsed float64
	var errCount int

	// print the challenge
	fmt.Printf("\n%s%s%s%s\n", colorPurple, challengeLimiter, colorReset, challenge)

	// build prompt and parse input
	fmt.Print(colorCyan, challengeLimiter, colorReset)

	start := time.Now()
	for _, char := range challenge {
		r := readInput(string(char))
		fmt.Print(r.char)

		// quit challenge when exit
		if r.quit {
			return len(challenge), float64(time.Since(start)) / float64(time.Second)
		}

		if r.failure {
			errCount += 1
		}
	}
	fmt.Println()

	elapsed = float64(time.Since(start)) / float64(time.Second)

	return errCount, elapsed
}

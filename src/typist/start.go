package typist

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Main Wrapper for a set of tests
func Start(challenge_file string, rate float64) {

	var c Challenges
	var average float64
	var errRate float64

	c.ReadFile(challenge_file)

	// Initialization of Keyboard
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	LogMessage(fmt.Sprintf("Welcome to typist v%s", version), "green")
	fmt.Printf("Your Challenge: %s\n", c.Description)

	// Loop over different challenges
	for _, i := range c.Lines {
		errCount, elapsed := ShowChallenge(i)

		// statistics
		errRate = float64(errCount) / float64(len(i)) * 100
		average += errRate

		if errRate < rate {
			LogMessage(fmt.Sprintf("%s %.2f%% (<%.2f%%) error rate in %.2fs\n", string(successChar), errRate, rate, elapsed), "green")
		} else {
			LogMessage(fmt.Sprintf("%s %.2f%% (<%.2f%%) error rate in %.2fs\n", string(failureChar), errRate, rate, elapsed), "red")
		}
	}

	fmt.Printf("\\o/ Average error rate: %.2f%%\n", average/float64(len(c.Lines)))

}

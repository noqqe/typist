// Package main provides a typing test
package main

import (
  "fmt"
  "log"
  "time"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "github.com/eiannone/keyboard"
)

const (
  errLimit float64 = 5 // Error Rate Limit
  challengeLimiter string = ">> "
  resultLimiter string = "== "
  colorRed string = "\033[31m"
  colorPurple string = "\033[35m"
  colorCyan string = "\033[36m"
  colorGreen string = "\033[32m"
  colorReset string = "\033[0m"
)

type Challenges struct {
    Lines []string `challenges`
}

// Main Loop
func main() {

  var c Challenges
  var average float64

  c.readFile("challenges.yml")

  // Initialization of Keyboard
  err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

  // Loop over different challenges
  for _, i := range c.Lines {
    succeeded, errRate, elapsed := challengeTypist(i)

    average += errRate

    if succeeded == true {
      fmt.Printf("SUCESS! %.2f%% (<%.2f%%) error rate in %.2fs\n", errRate, errLimit, elapsed)
    } else {
      fmt.Printf("FAILURE! %.2f%% (<%.2f%%) error rate in %.2fs\n", errRate, errLimit, elapsed)
    }
  }

  fmt.Printf("\n\\o/ Average error rate: %.2f%%\n", average/float64(len(c.Lines)))
}

// Read formatted yaml file
func (c *Challenges) readFile(path string) *Challenges {

    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        log.Printf("yamlFile.Get err #%v ", err)
    }

    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

// Read input from Keystrokes and compare with expected input
func readInput(expect string) (bool, string) {

  char, key, err := keyboard.GetKey()
  if (err != nil) {
      log.Fatalf("GetKey: %v", err)
  }

  // check if non alpha numeric input (like space)
  // fmt.Print("\n", int(key), int([]rune(expect)[0]), "\n")

  if char == 0 {
    if int(key) == int([]rune(expect)[0]) {
      return false, string(key)
    }
  }

  // fmt.Println(string(char), expect)
  if string(char) == expect {
    return false, string(char)
  } else {
    return true, fmt.Sprintf("%sx%s", colorRed, colorReset)
  }

}

// Prompt user the challenge and return if succeeded and what percentage
func challengeTypist(challenge string) (bool, float64, float64) {

  var elapsed float64
  var errCount float64
  var errRate float64

  // print the challenge
  fmt.Printf("\n%s%s%s%s\n", colorPurple, challengeLimiter, colorReset, challenge)

  // build prompt and parse input
  fmt.Print(colorCyan, challengeLimiter, colorReset)

  start := time.Now()
  for _, char := range challenge {
    failure, c := readInput(string(char))
    fmt.Print(c)
    if failure {
      errCount += 1
    }
  }
  fmt.Println()

  elapsed = float64(time.Since(start))/float64(time.Second)
  errRate = errCount/float64(len(challenge))*100

  if errRate < errLimit {
    return true, errRate, elapsed
  }

  return false, errRate, elapsed
}


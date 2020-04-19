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

type typeResponse struct {
  failure bool
  quit bool
  char string
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
func readInput(expect string) typeResponse {

  var resp typeResponse

  char, key, err := keyboard.GetKey()
  if (err != nil) {
      log.Fatalf("GetKey: %v", err)
  }

  // check if non alpha numeric input (like space)
  // fmt.Print("\n", int(key), int([]rune(expect)[0]), "\n")
  if char == 0 {

    // check user wants to quit
    if int(key) == 3 {
      resp.failure = false
      resp.quit = true
      resp.char = "abort...\n"
      return resp
    }

    // compare if real key like space
    if int(key) == int([]rune(expect)[0]) {
      resp.failure = false
      resp.quit = false
      resp.char = string(key)
      return resp
    }
  }

  // fmt.Println(string(char), expect)
  if string(char) == expect {
      resp.failure = false
      resp.quit = false
      resp.char = string(char)
      return resp
  } else {
      resp.failure = true
      resp.quit = false
      resp.char = fmt.Sprintf("%sx%s", colorRed, colorReset)
      return resp
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
    r := readInput(string(char))
    fmt.Print(r.char)

    // quit challenge when exit
    if r.quit {
      return false, 100, float64(time.Since(start))/float64(time.Second)
    }

    if r.failure {
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


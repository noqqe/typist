// Package main provides a typing test
package main

import (
  "fmt"
  "flag"
  "os"
  "time"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "github.com/eiannone/keyboard"
)

const (
  challengeLimiter string = ">> "
  colorCyan string = "\033[36m"
  colorGreen string = "\033[32m"
  colorPurple string = "\033[35m"
  colorRed string = "\033[31m"
  colorReset string = "\033[0m"
  failureChar = 10008 // cross heavy utf8
  successChar = 10004 // checkmark utf8
  version string = "0.0.2"
)

type Challenges struct {
    Description string `description`
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
  var rate float64
  var challenge_file string
  var errRate float64

  // flags declaration using flag package
  flag.Float64Var(&rate, "r", 5, "Allowed error rate. Default: 5")
  flag.StringVar(&challenge_file, "f", "challenges/intro-1.yml", "Specify challenge file")
  flag.Parse()

  c.readFile(challenge_file)

  // Initialization of Keyboard
  err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

  logMessage(fmt.Sprintf("Welcome to typist v%s", version), "green")
  fmt.Printf("Your Challenge: %s\n", c.Description)

  // Loop over different challenges
  for _, i := range c.Lines {
    errCount, elapsed := showChallenge(i)

    // statistics
    errRate = float64(errCount)/float64(len(i))*100
    average += errRate

    if errRate < rate {
      logMessage(fmt.Sprintf("%s %.2f%% (<%.2f%%) error rate in %.2fs\n", string(successChar), errRate, rate, elapsed), "green")
    } else {
      logMessage(fmt.Sprintf("%s %.2f%% (<%.2f%%) error rate in %.2fs\n", string(failureChar), errRate, rate, elapsed), "red")
    }
  }

  fmt.Printf("\\o/ Average error rate: %.2f%%\n", average/float64(len(c.Lines)))
}

func logMessage(message string, color string) {
  if (color == "red") {
    fmt.Printf("%s%s%s\n", colorRed, message, colorReset)
  } else if (color == "green") {
    fmt.Printf("%s%s%s\n", colorGreen, message, colorReset)
  } else {
    fmt.Printf("%s\n", message)
  }
}

// Read formatted yaml file
func (c *Challenges) readFile(path string) *Challenges {

    yamlFile, err := ioutil.ReadFile(path)
    if err != nil {
        logMessage("Could not open file", "red")
        os.Exit(1)
    }

    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        logMessage(fmt.Sprintf("Unmarshal %v", err), "red")
        os.Exit(1)
    }

    return c
}

// Read input from Keystrokes and compare with expected input
func readInput(expect string) typeResponse {

  var resp typeResponse
  resp.failure = false
  resp.quit = false


  char, key, err := keyboard.GetKey()
  if (err != nil) {
      logMessage(fmt.Sprintf("GetKey: %v ", err), "red")
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

// Prompt user the challenge and return if succeeded and what percentage
func showChallenge(challenge string) (int, float64) {

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
      return len(challenge), float64(time.Since(start))/float64(time.Second)
    }

    if r.failure {
      errCount += 1
    }
  }
  fmt.Println()

  elapsed = float64(time.Since(start))/float64(time.Second)

  return errCount, elapsed
}


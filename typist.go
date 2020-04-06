// Package main provides a typing test
package main

import (
  "fmt"
  "bufio"
  "os"
  "log"
  "time"
  "io/ioutil"
  "gopkg.in/yaml.v2"
  "strings"
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


  for _, i := range c.Lines {
    succeeded, errRate, elapsed := challengeTypist(i)

    average += errRate

    if succeeded {
      fmt.Printf("Success took %.2fs, errors rate: %.2f%%\n", elapsed, errRate)
    } else {
      fmt.Printf("Eh...nope took you %.2fs to fail miserbly, error rate too high: %.2f%%\n", elapsed, errRate)
    }
  }

  fmt.Printf("\\o/ Average error rate: %.2f%%\n", average/float64(len(c.Lines)))
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

// Prompt user the challenge and return if succeeded and what percentage
func challengeTypist(challenge string) (bool, float64, float64) {

  // print the challenge
  fmt.Printf("\n%s%s%s%s\n", colorPurple, challengeLimiter, colorReset, challenge)

  // build prompt and parse input
  fmt.Print(colorCyan, challengeLimiter, colorReset)
  reader := bufio.NewReader(os.Stdin)

  start := time.Now()
  input, err := reader.ReadString('\n')
  if err != nil {
    log.Fatal(err)
  }
  elapsed := float64(time.Since(start))/float64(time.Second)

  succeded, errRate := compare(challenge, input)

  return succeded, errRate, elapsed
}

// Compares input with challenge and calculates error rate in percent
func compare(challenge string, input string) (bool, float64) {

  var max float64 = float64(len(challenge))
  var errCount float64 = 0
  var challengeSucceded bool = false

  // pad input to match challenge if too short
  for i := len(input); len(input) < len(challenge); i++ {
    input += "\x00"
  }

  // run comparsion char by char
  fmt.Print(colorGreen, resultLimiter, colorReset)

  // loop over input in length of original challenge
  for pos, _ := range input[:len(challenge)] {
    if input[pos] == challenge[pos] {
      fmt.Printf("%s", string(challenge[pos]))
    } else {
      errCount += 1
      fmt.Printf("%sx%s", colorRed, colorReset)
    }
  }

  if len(input) > len(challenge) {
    lendiff := float64(len(input)-1-len(challenge))
    errCount += lendiff
    fmt.Printf("%s%s%s", colorRed, strings.Repeat("x", int(lendiff)) ,colorReset)
  }

  fmt.Print("\n")

  // build results
  var errRate float64 = errCount/max*100

  if errRate < errLimit {
    challengeSucceded = true
  }

  return challengeSucceded, errRate
}

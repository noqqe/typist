// Package main provides a typing test
package main

import (
  "fmt"
  "bufio"
  "os"
  "log"
  "time"
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


func main() {

  // some challenges, just for now to test
  challenges := []string{
    "echo \"foo bar\"\n",
    "grep bar apache.log\n",
    "sed -i 's/apache/nginx/g'\n",
  }

  for i := range challenges {
    succeeded, errRate, elapsed := challengeTypist(challenges[i])

    if succeeded {
      fmt.Printf("Success took %.2fs, errors rate: %.2f%%\n", elapsed, errRate)
    } else {
      fmt.Printf("Eh...nope took you %.2fs to fail miserbly, error rate too high: %.2f%%\n", elapsed, errRate)
    }
  }

}

// Prompt user the challenge and return if succeeded and what percentage
func challengeTypist(challenge string) (bool, float64, float64) {
// print the challenge
  fmt.Printf("\n%s%s%s%s", colorPurple, challengeLimiter, colorReset, challenge)

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
    input += " "
  }

  // run comparsion char by char
  fmt.Print(colorGreen, resultLimiter, colorReset)

  // loop over input in length of original challenge
  for pos, _ := range input[:len(challenge)-1] {
    if input[pos] == challenge[pos] {
      fmt.Printf("%s", string(challenge[pos]))
    } else {
      errCount += 1
      fmt.Printf("%sx%s", colorRed, colorReset)
    }
  }
  fmt.Print("\n")

  // build results
  var errRate float64 = errCount/max*100

  if errRate < errLimit {
    challengeSucceded = true
  }

  return challengeSucceded, errRate
}

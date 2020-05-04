// Package main provides a typing test
package main

import (
	"flag"

	"github.com/noqqe/typist/src/typist"
)

// Main Loop
func main() {

	var rate float64
	var challenge_file string

	// flags declaration using flag package
	flag.Float64Var(&rate, "r", 5, "Allowed error rate. Default: 5")
	flag.StringVar(&challenge_file, "f", "challenges/intro-1.yml", "Specify challenge file")
	flag.Parse()

	typist.Start(challenge_file, rate)
}

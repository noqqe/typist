package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()
	for {
    k()
	}
}

func k() {

	fmt.Println("Press ESC to quit")
  char, key, err := keyboard.GetKey()
  if (err != nil) {
    panic(err)
  } else if (key == keyboard.KeyEsc) {
    return
  }
  fmt.Printf("You pressed: %q\r\n", char)
}

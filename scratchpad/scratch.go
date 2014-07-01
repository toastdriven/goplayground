package main

import (
	"fmt"
	// "math/rand"
)

func main() {
	foo := []string{"H", "e", "l", "l", "o"}
	fmt.Println(foo)
	foo = append(foo[:2], foo[3:]...)
	fmt.Println(foo)
}

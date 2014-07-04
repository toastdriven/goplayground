package main

import (
	"bufio"
	"fmt"
	"os"
)

type Dictionary struct {
	Words map[string]int
}

func NewDictionary() Dictionary {
	dict := Dictionary{}
	dict.Words = make(map[string]int)
	return dict
}

func (dict *Dictionary)Add(word string) {
	_, ok := dict.Words[word]

	if ok {
		dict.Words[word]++
	} else {
		dict.Words[word] = 1
	}
}

func main() {
	dict := NewDictionary()

	// Collect the input.
	fmt.Println("Dump in some text, followed by a blank line please!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		dict.Add(scanner.Text())
	}

	for word, count := range dict.Words {
		fmt.Println(word, "seen:", count)
	}
}

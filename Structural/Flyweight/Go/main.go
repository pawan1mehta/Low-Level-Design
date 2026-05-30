package main

import "fmt"

type TextCharacter interface {
	Display(size int)
}

type Letter struct {
	ch byte
}

func NewLetter(char byte) *Letter {
	fmt.Printf("Creating new flyweight for: '%c' \n", char)
	return &Letter{
		ch: char,
	}
}

func (l *Letter) Display(size int) {
	fmt.Printf("Character: %c with font size: %d \n", l.ch, size)
}

var characters = make(map[byte]*Letter)

func GetLetter(char byte) *Letter {
	if _, ok := characters[char]; !ok {
		characters[char] = NewLetter(char)
	} else {
		fmt.Printf("Reusing flyweight for: '%c' \n", char)
	}
	return characters[char]
}

func main() {
	doc := "Pawan Mehta"

	for _, char := range doc {
		letter := GetLetter(byte(char))
		letter.Display(10)
	}
}

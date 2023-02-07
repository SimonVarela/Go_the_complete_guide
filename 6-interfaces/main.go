package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Couldn't open file:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}

/*
type shape interface {
	getArea() float64
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

type triangle struct {
	height float64
	base   float64
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func main() {
	square := square{4}
	triangle := triangle{2, 2}
	fmt.Print("Square area: ")
	printArea(square)
	fmt.Print("Triangle area: ")
	printArea(triangle)
}
*/
/*
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
*/
/*
import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	spanish := spanishBot{}
	english := englishBot{}

	printGreeting(english)
	printGreeting(spanish)
}

func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Epa chamo!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
*/

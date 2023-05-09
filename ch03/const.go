package ch03

import "fmt"

const (
	Apple = iota
	Grape
	Orange
)

func PrintConst() {
	fmt.Println(Apple)
	fmt.Println(Grape)
	fmt.Println(Orange)
}

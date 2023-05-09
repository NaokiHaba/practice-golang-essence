package ch03

import "fmt"

// Speaker は鳴くことができるもののインターフェース
type Speaker interface {
	Speak() error
}

// Dog は犬の構造体
type Dog struct{}

// Speak は犬が鳴く
func (d *Dog) Speak() error {
	fmt.Println("わんわん")
	return nil
}

// Cat は猫の構造体
type Cat struct{}

// Speak は猫が鳴く
func (c *Cat) Speak() error {
	fmt.Println("にゃーにゃー")
	return nil
}

// DoSpeak は鳴くことができるものを鳴かせる
func DoSpeak(s Speaker) error {
	return s.Speak()
}

func PrintSpeaker() {
	dog := &Dog{}
	cat := &Cat{}

	DoSpeak(dog) // わんわん
	DoSpeak(cat) // にゃーにゃー
}

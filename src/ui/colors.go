package ui

import "fmt"

const (
	Cyan  = "\033[36m"
	Reset = "\033[0m"
	Green = "\033[32m"
	Red   = "\033[31m"
)

func SuccessText(text string) {
	fmt.Printf("%s%s%s\n", Green, text, Reset)
}

func ErrorText(text string) {
	fmt.Printf("%s%s%s\n", Red, text, Reset)
}

func InfoText(text string) {
	fmt.Printf("%s%s%s\n", Cyan, text, Reset)
}
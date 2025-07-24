package ui

import "fmt"

const (
	Cyan  string = "\033[36m"
	Reset string = "\033[0m"
	Green string = "\033[32m"
	Red   string = "\033[31m"
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
package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Question struct {
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name string
	Score int
	Questions []Question
}

var Cyan = "\033[36m"
var Reset = "\033[0m"
var Green = "\033[32m"
var Red = "\033[31m"

func (g *GameState) Init() {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Print("Escreva seu nome: ")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler o nome")
	}

	g.Name = name[:len(name)-2]
	
	fmt.Printf("Olá, %s! Vamos começar o quiz.\n", g.Name)
}

func (g *GameState) ProcessCSV() {
	file, err := os.Open("quizgo.csv")
	if err != nil {
		panic("Erro ao abrir o arquivo CSV")
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler o arquivo CSV")
	}

	for index, record := range records {
		if index > 0 {
			correctAnswer, _ := toInt(record[len(record)-1])
			question := Question {
				Text: record[0],
				Answer: correctAnswer,
			}

			for i := 1; i < len(record)-1; i++ {
				question.Options = append(question.Options, record[i])
			}
			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {
	reader := bufio.NewReader(os.Stdin)

	for index, question := range g.Questions {
		fmt.Printf("%s%d. %s\n", Cyan, index+1, question.Text)
		fmt.Print(Reset)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Print("Escolha uma opção: ")

		var answer int
		var err error
		for {
			read, _ := reader.ReadString('\n')
			answer, err = toInt(read[:len(read)-2])
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			g.Score++
			fmt.Printf("%sResposta correta!%s\n", Green, Reset)
		} else {
			fmt.Printf("%sResposta incorreta!%s\n", Red, Reset)
		}
	}
}

func main() {
	game := &GameState{}
	go game.ProcessCSV()
	game.Init()
	game.Run()
}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("não é permitido caracteres não numéricos")
	}
	return i, nil
}
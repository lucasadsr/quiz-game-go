package main

import (
	"bufio"
	"encoding/csv"
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
			question := Question {
				Text: record[0],
				Answer: toInt(record[len(record)-1]),
			}

			for i := 1; i < len(record)-1; i++ {
				question.Options = append(question.Options, record[i])
			}
			g.Questions = append(g.Questions, question)
		}
	}
}

func main() {
	game := &GameState{}
	game.Init()
	go game.ProcessCSV()
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("Erro ao converter string para inteiro: " + err.Error())
	}
	return i
}
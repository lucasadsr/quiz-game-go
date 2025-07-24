package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/lucasadsr/quiz-game-go/src/ui"
	"github.com/lucasadsr/quiz-game-go/src/utils"
)

type Question struct {
	Text string
	Options []string
	Answer int
}

type GameState struct {
	Name     string
	Score    int
	Subject  string
	Questions []Question
}

var Subjects = []string{"Português", "História", "Geografia", "Ciências"}

func (g *GameState) Init() {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Print("Escreva seu nome: ")
	reader := bufio.NewReader(os.Stdin)

	var name string
	var err error
	for {
		name, err = reader.ReadString('\n')
		if err != nil || len(strings.TrimSpace(name)) < 2 {
			fmt.Println("Digite um nome válido.")
			continue
		}
		break
	}

	g.Name = name[:len(name)-2]

	fmt.Printf("Olá, %s! Vamos começar o quiz.\n", g.Name)
}

func (g *GameState) SetSubject() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Qual matéria você gostaria de estudar?")
	for i, subject := range Subjects {
		fmt.Printf("[%d] %s\n", i+1, subject)
	}

	fmt.Print("Escolha uma opção: ")

	var subject int
	var err error
	for {
		read, _ := reader.ReadString('\n')
		subject, err = utils.ToInt(read[:len(read)-2])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		g.Subject = Subjects[subject-1]
		fmt.Printf("Você escolheu estudar %s.\n", g.Subject)
		break
	}
	fmt.Println(strings.Repeat("-", 30))
}

func (g *GameState) ProcessCSV() {
	filename := fmt.Sprintf("./questions/questions-%s.csv", g.Subject)
	file, err := os.Open(filename)
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
			correctAnswer, _ := utils.ToInt(record[len(record)-1])
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
		ui.InfoText(fmt.Sprintf("%d. %s", index+1, question.Text))

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Print("Escolha uma opção: ")

		var answer int
		var err error
		for {
			read, _ := reader.ReadString('\n')
			answer, err = utils.ToInt(read[:len(read)-2])
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		if answer == question.Answer {
			g.Score++
			ui.SuccessText("Resposta correta!")
		} else {
			ui.ErrorText("Resposta incorreta!")
		}
		fmt.Println(strings.Repeat("-", 30))
	}
}

func main() {
	game := &GameState{}
	game.Init()
	game.SetSubject()
	game.ProcessCSV()
	game.Run()

	fmt.Printf("Fim do quiz, %s! Sua pontuação final é: %d/%d\n", game.Name, game.Score, len(game.Questions))
}
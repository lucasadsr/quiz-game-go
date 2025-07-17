package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	game := &GameState{}
	game.Init()
}
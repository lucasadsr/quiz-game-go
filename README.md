# Quiz Go 🎯

Um jogo de quiz interativo desenvolvido em Go, onde o usuário responde perguntas de múltipla escolha carregadas de um arquivo CSV.

## 📋 Descrição

O Quiz Go é uma aplicação de linha de comando que permite aos usuários participarem de um quiz personalizado. O programa carrega perguntas de um arquivo CSV, apresenta as opções de resposta e calcula a pontuação final do jogador. Com uma interface colorida e amigável, o jogo oferece uma experiência interativa e educativa.

## 🚀 Tecnologias Utilizadas

- **Go 1.24.3** - Linguagem de programação principal
- **Bibliotecas padrão do Go:**
  - `bufio` - Para leitura de entrada do usuário
  - `encoding/csv` - Para processamento de arquivos CSV
  - `fmt` - Para formatação de saída
  - `os` - Para operações do sistema operacional
  - `strconv` - Para conversão de strings
  - `strings` - Para manipulação de strings

## 🏗️ Arquitetura

O projeto segue uma arquitetura simples e orientada a objetos:

### Estruturas Principais

- **`Question`** - Representa uma pergunta do quiz com:

  - Texto da pergunta
  - Lista de opções
  - Índice da resposta correta

- **`GameState`** - Gerencia o estado do jogo contendo:
  - Nome do jogador
  - Pontuação atual
  - Lista de perguntas carregadas

### Métodos Principais

- `Init()` - Inicializa o jogo e captura o nome do usuário
- `ProcessCSV()` - Carrega e processa as perguntas do arquivo CSV
- `Run()` - Executa o loop principal do quiz
- `toInt()` - Função auxiliar para conversão e validação de entrada

### Características Técnicas

- **Processamento concorrente**: O carregamento do CSV é executado em uma goroutine
- **Validação de entrada**: Tratamento de erros para entradas inválidas
- **Interface colorida**: Uso de códigos ANSI para colorização do terminal
- **Tratamento de arquivos**: Leitura segura de arquivos CSV com defer para fechamento

## 📁 Estrutura do Projeto

```
quiz-game/
├── main.go          # Código principal da aplicação
├── quizgo.csv       # Arquivo com as perguntas do quiz
├── go.mod           # Arquivo de módulo do Go
└── README.md        # Documentação do projeto
```

## ⚙️ Como Executar

### Pré-requisitos

- Go 1.24.3 ou superior instalado
- Terminal/Prompt de comando

### Passos para execução

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/lucasadsr/quiz-game-go.git
   cd quiz-game-go
   ```

2. **Execute o programa:**

   ```bash
   go run main.go
   ```

3. **Siga as instruções no terminal:**
   - Digite seu nome quando solicitado
   - Responda as perguntas escolhendo o número da opção desejada
   - Veja sua pontuação final ao concluir o quiz

### Alternativa - Compilar e executar

```bash
# Compilar o programa
go build -o quiz-game main.go

# Executar o binário (Windows)
quiz-game.exe

# Executar o binário (Linux/Mac)
./quiz-game
```

## 📝 Formato do Arquivo CSV

O arquivo `quizgo.csv` deve seguir o formato:

```csv
Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta Correta
Qual é o maior planeta do sistema solar?,Terra,Marte,Júpiter,Saturno,3
```

- **Primeira linha**: Cabeçalho (será ignorado)
- **Colunas subsequentes**:
  - Texto da pergunta
  - Opções de resposta (pode ter quantas quiser)
  - Última coluna: Número da resposta correta (baseado em 1)

## 🎮 Como Jogar

1. Execute o programa
2. Digite seu nome quando solicitado
3. Para cada pergunta:
   - Leia a pergunta apresentada
   - Escolha uma das opções numeradas
   - Digite o número da sua resposta
4. Ao final, veja sua pontuação total

## 🌟 Funcionalidades

- ✅ Interface colorida no terminal
- ✅ Carregamento dinâmico de perguntas via CSV
- ✅ Validação de entrada do usuário
- ✅ Pontuação em tempo real
- ✅ Feedback imediato para respostas
- ✅ Processamento concorrente para melhor performance

## 🔧 Personalização

Para adicionar suas próprias perguntas:

1. Edite o arquivo `quizgo.csv`
2. Adicione novas linhas seguindo o formato especificado
3. Execute o programa normalmente

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## 👨‍💻 Autor

**Lucas Ribeiro** - [@lucasadsr](https://github.com/lucasadsr)

---

⭐ Se você gostou do projeto, não esqueça de dar uma estrela no repositório!

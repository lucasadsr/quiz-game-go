# Quiz Go 🎯

Um jogo de quiz interativo desenvolvido em Go, onde o usuário pode escolher entre diferentes matérias e responder perguntas de múltipla escolha carregadas de arquivos CSV específicos.

## 📋 Descrição

O Quiz Go é uma aplicação de linha de comando que permite aos usuários participarem de um quiz educativo com múltiplas categorias. O programa oferece a escolha entre diferentes matérias, carrega perguntas específicas da matéria escolhida, apresenta as opções de resposta e calcula a pontuação final do jogador. Com uma interface colorida e amigável, o jogo oferece uma experiência interativa e educativa personalizada.

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
  - Matéria escolhida
  - Lista de perguntas carregadas

### Métodos Principais

- `Init()` - Inicializa o jogo e captura o nome do usuário
- `SetSubject()` - Permite ao usuário escolher a matéria do quiz
- `ProcessCSV()` - Carrega e processa as perguntas do arquivo CSV da matéria selecionada
- `Run()` - Executa o loop principal do quiz
- `toInt()` - Função auxiliar para conversão e validação de entrada

### Características Técnicas

- **Sistema de categorias**: Quatro matérias disponíveis (Português, História, Geografia, Ciências)
- **Carregamento dinâmico**: Perguntas carregadas conforme a matéria escolhida
- **Validação de entrada**: Tratamento de erros para entradas inválidas
- **Interface colorida**: Uso de códigos ANSI para colorização do terminal
- **Tratamento de arquivos**: Leitura segura de arquivos CSV com defer para fechamento

## 📁 Estrutura do Projeto

```
quiz-game/
├── main.go                    # Código principal da aplicação
├── questions/                 # Diretório com perguntas por matéria
│   ├── questions-Português.csv    # Perguntas de Português
│   ├── questions-História.csv     # Perguntas de História
│   ├── questions-Geografia.csv    # Perguntas de Geografia
│   └── questions-Ciências.csv     # Perguntas de Ciências
├── go.mod                     # Arquivo de módulo do Go
└── README.md                  # Documentação do projeto
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
   - Escolha a matéria que deseja estudar (Português, História, Geografia ou Ciências)
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

## 📝 Formato dos Arquivos CSV

Os arquivos de perguntas estão organizados por matéria na pasta `questions/` e devem seguir o formato:

```csv
Pergunta,Opção 1,Opção 2,Opção 3,Opção 4,Resposta Correta
Qual das palavras abaixo está escrita corretamente?,Exessão,Excessão,Exceção,Eceção,2
```

### Matérias Disponíveis:

- **questions-Português.csv** - Perguntas sobre gramática, literatura e redação
- **questions-História.csv** - Perguntas sobre história geral e do Brasil
- **questions-Geografia.csv** - Perguntas sobre geografia física e humana
- **questions-Ciências.csv** - Perguntas sobre biologia, física e química

### Formato das colunas:

- **Primeira linha**: Cabeçalho (será ignorado)
- **Colunas subsequentes**:
  - Texto da pergunta
  - Opções de resposta (pode ter quantas quiser)
  - Última coluna: Número da resposta correta (baseado em 1)## 🎮 Como Jogar

1. Execute o programa
2. Digite seu nome quando solicitado
3. Escolha a matéria que deseja estudar:
   - [1] Português
   - [2] História
   - [3] Geografia
   - [4] Ciências
4. Para cada pergunta:
   - Leia a pergunta apresentada
   - Escolha uma das opções numeradas
   - Digite o número da sua resposta
5. Ao final, veja sua pontuação total

## 🌟 Funcionalidades

- ✅ Interface colorida no terminal
- ✅ Sistema de categorias por matéria
- ✅ Carregamento dinâmico de perguntas por matéria escolhida
- ✅ Quatro matérias disponíveis (Português, História, Geografia, Ciências)
- ✅ Validação de entrada do usuário
- ✅ Pontuação em tempo real
- ✅ Feedback imediato para respostas
- ✅ Organização modular dos arquivos de perguntas

## 🔧 Personalização

### Adicionando novas perguntas:

1. Edite o arquivo CSV da matéria desejada na pasta `questions/`
2. Adicione novas linhas seguindo o formato especificado
3. Execute o programa normalmente

### Adicionando uma nova matéria:

1. Crie um novo arquivo CSV seguindo o padrão: `questions-[NomeDaMatéria].csv`
2. Adicione a nova matéria no array `Subjects` no arquivo `main.go`:
   ```go
   var Subjects = []string{"Português", "História", "Geografia", "Ciências", "Nova Matéria"}
   ```
3. Compile e execute o programa

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.

## 👨‍💻 Autor

**Lucas Ribeiro** - [@lucasadsr](https://github.com/lucasadsr)

---

⭐ Se você gostou do projeto, não esqueça de dar uma estrela no repositório!

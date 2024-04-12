package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var fileScanner *bufio.Scanner
var questions []Question
var visitedQuestions []int

func nextLine() string {
	fileScanner.Scan()
	return fileScanner.Text()
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	var r = rand.New(rand.NewSource(time.Now().UnixNano()))

	stdin := bufio.NewReader(os.Stdin)
	file, err := os.Open("questions.txt")

	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner = bufio.NewScanner(file)

	visitedQuestions = []int{}

	for fileScanner.Scan() {
		pergunta := Question{
			Question:    nextLine(),
			Op1:         nextLine(),
			Op2:         nextLine(),
			Op3:         nextLine(),
			Op4:         nextLine(),
			RightAnswer: nextLine(),
		}

		questions = append(questions, pergunta)
	}

	premiacoes := [11]int{500, 1000, 2500, 5000, 10000, 15000, 25000, 50000, 100000, 500000, 10000000}
	fmt.Printf("Seja bem vindo ao Jogo do Calebão!\n")
	fmt.Printf("Deseja começar?\n 1- Sim\n 2- Não\n ")

	var comecar string
	var comecou bool

	fmt.Scanf("%s", &comecar)
	fmt.Printf("O jogo é composto por 11 perguntas se acertar todas você leva para casa o incrivel premio de 1 milhao de reais\nAs questions vão vir na formatação a seguir: \nEnunciado da Questão\nOpção 1\nOpção 2\nOPção 3\nOpção 4\n")
	if comecar == "1" {
		comecou = true
	} else {
		fmt.Printf(" Volte Sempre!\n")
		comecou = false
	}
	if comecou == true {
		premio := 0
		cartasusadas := 0
		universitariosusados := 0
		pularusados := 0

		for {
			var comando string
			fmt.Scanf("%s", &comando)

			for comando != "Q" {
				fmt.Printf("Valendo R$ %d!\n", premiacoes[premio])

				var numeroPergunta int

				for {
					// fmt.Printf(".\n")

					numeroPergunta = r.Intn(len(questions) - 1)

					// fmt.Printf("%v\n", contains(visitedQuestions, numeroPergunta))
					// fmt.Printf("%v\n", visitedQuestions)
					// fmt.Printf("%v\n", numeroPergunta)
					// fmt.Printf("%v\n", len(questions))

					if !contains(visitedQuestions, numeroPergunta) {
						break
					}

					if len(visitedQuestions) == len(questions) {
						fmt.Printf("Parabéns você zerou o jogo!!!")
						return
					}
				}

				visitedQuestions = append(visitedQuestions, numeroPergunta)
				pergunta := questions[numeroPergunta]

				repetir := true
				cartas := 0

				for repetir {
					repetir = false

					fmt.Printf("%s\n\n", pergunta.Question)

					if cartas == 0 || "1" == pergunta.RightAnswer {
						fmt.Printf("1 - %s\n", pergunta.Op1)
					} else if "1" != pergunta.RightAnswer {
						fmt.Printf("\n")
						cartas--
					}
					if cartas == 0 || "2" == pergunta.RightAnswer {
						fmt.Printf("2 - %s\n", pergunta.Op2)
					} else if "2" != pergunta.RightAnswer {
						fmt.Printf("\n")
						cartas--
					}
					if cartas == 0 || "3" == pergunta.RightAnswer {
						fmt.Printf("3 - %s\n", pergunta.Op3)
					} else if "3" != pergunta.RightAnswer {
						fmt.Printf("\n")
						cartas--
					}
					if cartas == 0 || "4" == pergunta.RightAnswer {
						fmt.Printf("4 - %s\n", pergunta.Op4)
					} else if "4" != pergunta.RightAnswer {
						fmt.Printf("\n")
						cartas--
					}

					for {
						_, err := fmt.Scanf("%s\n", &comando)
						if err == nil {
							break
						}

						stdin.ReadString('\n')
						//fmt.Println("Sorry, invalid input. Please enter an integer: ")
					}

					if comando == "C" || comando == "c" {
						if cartasusadas == 2 {
							fmt.Printf("Você não tem mais cartas\n")
							repetir = true
							continue
						}
						fmt.Printf("Escolha uma carta (número de 1 a 4)\n")
						fmt.Scanf("%s\n", &comando)
						if comando < "0" || comando > "4" {
							fmt.Printf("Carta inválida :(\nDigite outra por favor.\n")
							fmt.Scanf("%s\n", &comando)
						}
						cartas = rand.Intn(3)
						fmt.Printf("Vamos tirar %d opções...\n", cartas)
						repetir = true
						cartasusadas++
						continue
					}
					if comando == "U" || comando == "u" && universitariosusados < 2 {
						uni1 := rand.Intn(90) + 10
						fmt.Printf("%d%%, dos universitarios acham que é %s\n", uni1, pergunta.RightAnswer)
						universitariosusados++
					}
					if comando == "P" || comando == "p" && pularusados < 2 {
						pularusados++
						continue
					}

					if comando > "4" || comando < "0" {
						fmt.Printf("Resposta Inválida, Digite outro número por favor\n")
						fmt.Scanf("%s", &comando)
					}

					if comando != pergunta.RightAnswer {
						comando = "Q"

						fmt.Printf("Game Over!\n")

						if premio >= 2 {
							fmt.Printf("Você ganhou %d reais\n", premiacoes[premio-2])
						} else {
							fmt.Printf("Você não ganhou nada!\n")
						}

						premio = 0
						break
					} else {
						fmt.Printf("Correto!\n")
						premio++
					}

					if premio == 11 {
						fmt.Printf("Parabéns você é o grande vencedor!\n")
						break
					}
				}

			}
		}
	}
}

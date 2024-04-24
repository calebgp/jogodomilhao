package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/calebgp/jogodomilhao/models"
	"google.golang.org/api/option"
	"log"
	"math/rand"
	"time"
)

var (
	questions        []models.Question
	visitedQuestions []int
	premiacoes       = [11]int{500, 1000, 2500, 5000, 10000, 15000, 25000, 50000, 100000, 500000, 10000000}
	welcomeMessage   = "" +
		"O jogo é composto por 11 perguntas se acertar todas você leva para casa o incrivel premio de 1 milhao de reais\n" +
		"As questions vão vir na formatação a seguir: \n" +
		"Enunciado da Questão\n" +
		"Opção 1\n" +
		"Opção 2\n" +
		"Opção 3\n" +
		"Opção 4\n" +
		"Se precisar de ajuda para resolver qualquer uma das questões:\n" +
		"Cartas: Digite C\n" +
		"Universitátios: Digite U\n" +
		"Pular: Digite P\n\n"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func initializeQuestions() ([]models.Question, []int) {
	// Retrieve questions from database
	var questions []models.Question
	ctx := context.Background()
	sa := option.WithCredentialsFile("credentials.json")
	fmt.Println("Conectando a base de perguntas...")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
	colRef := client.Collection("questions").Documents(ctx)
	defer colRef.Stop() // add this line to ensure resources cleaned up
	fmt.Println("Buscando perguntas...")
	for {
		doc, err := colRef.Next()
		if err != nil {
			break
		}
		var q models.Question
		if err := doc.DataTo(&q); err != nil {
			log.Fatalf("doc.DataTo: %v", err)
		}
		questions = append(questions, q)
	}
	// Initialize visitedQuestions array
	var visitedQuestions []int

	return questions, visitedQuestions
}
func selectRandomQuestion(questions []models.Question, visitedQuestions []int) models.Question {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		questionNumber := r.Intn(len(questions) - 1)
		if !contains(visitedQuestions, questionNumber) {
			visitedQuestions = append(visitedQuestions, questionNumber)
			return questions[questionNumber]
		}

		if len(visitedQuestions) == len(questions) {
			fmt.Printf("Parabéns você zerou o jogo!!!")
			break
		}
	}

	return models.Question{}
}
func handleUserInput(question models.Question) (answer string) {
	var command string
	usedCards := 0
	usedCollegeStudents := 0
	usedSkips := 0

	for {
		command = ""

		for command != "Q" {

			repetir := true
			cartas := 0

			for repetir {
				repetir = false

				fmt.Printf("%s\n\n", question.Question)

				if cartas == 0 || "1" == question.RightAnswer {
					fmt.Printf("1 - %s\n", question.Op1)
				} else if "1" != question.RightAnswer {
					fmt.Printf("\n")
					cartas--
				}
				if cartas == 0 || "2" == question.RightAnswer {
					fmt.Printf("2 - %s\n", question.Op2)
				} else if "2" != question.RightAnswer {
					fmt.Printf("\n")
					cartas--
				}
				if cartas == 0 || "3" == question.RightAnswer {
					fmt.Printf("3 - %s\n", question.Op3)
				} else if "3" != question.RightAnswer {
					fmt.Printf("\n")
					cartas--
				}
				if cartas == 0 || "4" == question.RightAnswer {
					fmt.Printf("4 - %s\n", question.Op4)
				} else if "4" != question.RightAnswer {
					fmt.Printf("\n")
					cartas--
				}

				for {
					_, err := fmt.Scanf("%s\n", &command)
					if err == nil {
						break
					}

					//stdin.ReadString('\n')
					//fmt.Println("Sorry, invalid input. Please enter an integer: ")
				}

				if command == "C" || command == "c" {
					if usedCards == 2 {
						fmt.Printf("Você não tem mais cartas\n")
						repetir = true
						continue
					}
					fmt.Printf("Escolha uma carta (número de 1 a 4)\n")
					_, err := fmt.Scanf("%s\n", &command)
					if err != nil {
						return ""
					}
					if command < "0" || command > "4" {
						fmt.Printf("Carta inválida :(\nDigite outra por favor.\n")
						_, err := fmt.Scanf("%s\n", &command)
						if err != nil {
							return ""
						}

					}
					cartas = rand.Intn(3)
					fmt.Printf("Vamos tirar %d opções...\n", cartas)
					repetir = true
					usedCards++
					continue
				}
				if command == "U" || command == "u" && usedCollegeStudents < 2 {
					uni1 := rand.Intn(90) + 10
					fmt.Printf("%d%%, dos universitarios acham que é %s\n", uni1, question.RightAnswer)
					usedCollegeStudents++
				}
				if command == "P" || command == "p" && usedSkips < 2 {
					usedSkips++
					continue
				}

				if command > "4" || command < "0" {
					fmt.Printf("Resposta Inválida, Digite outro número por favor\n")
					_, err := fmt.Scanf("%s", &command)
					if err != nil {
						return ""
					}
				} else {
					return command
				}

			}

		}
	}

}
func playGame(questions []models.Question, visitedQuestions []int) {
	fmt.Printf(welcomeMessage)
	prize := 0

	for {
		fmt.Printf("Valendo R$ %d!\n", premiacoes[prize])

		// Select a random question that hasn't been asked yet

		question := selectRandomQuestion(questions, visitedQuestions)

		repetir := true

		for repetir {
			repetir = false
			// Handle user input
			answer := handleUserInput(question)

			// Check if user's answer is correct
			if answer == question.RightAnswer {
				fmt.Printf("Correto!\n")
				prize++
				visitedQuestions = append(visitedQuestions)

				if prize == 11 {
					fmt.Printf("Parabéns você é o grande vencedor!\n")
					break
				}
			} else {
				fmt.Printf("Game Over!\n")

				if prize >= 2 {
					fmt.Printf("Você ganhou %d reais\n", premiacoes[prize-2])
				} else {
					fmt.Printf("Você não ganhou nada!\n")
				}

				prize = 0
				return
			}
		}
	}
}

func main() {
	// Initialize game state and data
	questions, visitedQuestions = initializeQuestions()

	// Welcome message and game introduction
	fmt.Printf("Seja bem vindo ao Jogo do Calebão!\n")
	fmt.Printf("Deseja começar?\n 1- Sim\n 2- Não\n ")

	var comecar string
	_, err := fmt.Scanf("%s", &comecar)
	if err != nil {
		return
	}
	if comecar != "1" {
		fmt.Printf(" Volte Sempre!\n")
		return
	}
	playGame(questions, visitedQuestions)
}

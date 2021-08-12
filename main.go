package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var fileScanner *bufio.Scanner

func proximaLinha() string {
	fileScanner.Scan()
	return fileScanner.Text()
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	file, err := os.Open("perguntas.txt")
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner = bufio.NewScanner(file)

	premiacoes := [11]int{500, 1000, 2500, 5000, 10000, 25000, 50000, 100000, 200000, 500000, 1000000}
	fmt.Printf("Seja bem vindo ao Jogo do Calebão!\n")
	fmt.Printf("Deseja começar?\n 1- Sim\n 2- Não\n ")
	var comecar string
	var comecou bool
	fmt.Scanf("%s", &comecar)
	if comecar == "1" {
		comecou = true
	} else {
		fmt.Printf(" Volte Sempre!\n")
		comecou = false
	}
	if comecou == true {
		var pergunta string
		var opcao1, opcao2, opcao3, opcao4, gabarito string
		premio := 0
		for {
			var comando string
			fmt.Scanf("%s", &comando)
			for comando != "Q" {
				fmt.Printf("Valendo R$ %d!\n", premiacoes[premio])
				pergunta = proximaLinha()
				opcao1 = proximaLinha()
				opcao2 = proximaLinha()
				opcao3 = proximaLinha()
				opcao4 = proximaLinha()
				gabarito = proximaLinha()
				proximaLinha()
				fmt.Printf("%s\n\n", pergunta)
				fmt.Printf("1 - %s\n", opcao1)
				fmt.Printf("2 - %s\n", opcao2)
				fmt.Printf("3 - %s\n", opcao3)
				fmt.Printf("4 - %s\n", opcao4)
				for {
					_, err := fmt.Scanf("%s\n", &comando)
					if err == nil {
						break
					}

					stdin.ReadString('\n')
					fmt.Println("Sorry, invalid input. Please enter an integer: ")
				}

				if comando == "C" {
					fmt.Printf("Escolha uma carta (número de 1 a 4)\n")
					fmt.Scanf("%s", &comando)
					rand.Intn(4)
				}
				if comando == "U" {

				}
				if comando != gabarito {
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

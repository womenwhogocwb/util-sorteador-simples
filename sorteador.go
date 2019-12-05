package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Alterar participantes de acordo com o número necessário
	participantes := 30
	//Alterar segundosContagem de acordo com tempo desejado
	segundosContagem := 3

	nestante := time.Now()
	rand.Seed(nestante.UnixNano())
	random := 1 + rand.Intn(participantes)

	linhas()
	fmt.Println(nestante.Format("02-Jan-06 15:04:05"))
	linhas()

	contagem(segundosContagem)
	resultado(random)

}

func linhas() {
	for i := 1; i < 12; i++ {
		fmt.Print("=")
		sleepMenor()
	}
	fmt.Println("=")
	sleepMaior()
}

func contagem(segundos int) {
	frase := "🚀 Sorteando em:"

	for _, caracter := range frase {
		sleepMenor()
		fmt.Print(string(caracter))
	}

	for i := segundos; i > 0; i-- {
		fmt.Printf(" %d", i)
		for j := 0; j < 3; j++ {
			sleepMenor()
			fmt.Print(".")
		}
		sleepMaior()
	}
	fmt.Println("\n\n🍀 🍀 🍀 🍀 🍀 🍀 🍀 🍀")
	sleepMaior()
}

func resultado(numero int) {
	fmt.Println("O número sorteado foi:")
	sleepMaior()
	fmt.Printf(">>> %d <<<\n", numero)
	fmt.Println("============")
}

func sleepMaior() {
	time.Sleep(1 * time.Second)
}

func sleepMenor() {
	time.Sleep(100 * time.Millisecond)
}
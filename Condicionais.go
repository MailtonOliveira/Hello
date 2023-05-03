package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	heistReady := true

	if heistReady {
		fmt.Println("Vamos Lá")
	} else {
		fmt.Println("ação normal")
	}
	lockCombo := "3-1-25"
	robAttempt := "3-1-25"

	if lockCombo == robAttempt {
		fmt.Println("Cofre foi aberto")
	} else {
		fmt.Println("Errou otário")
	}
	vaultAmt := 2356468

	if vaultAmt >= 200000 {
		fmt.Println("Vamos Precisar de mais malas")
	}
	rightTime := true
	rightPlace := true

	if rightTime && rightPlace {
		fmt.Println("Saímos daqui")
	} else {
		fmt.Println("Seja paciente...")
	}
	enoughRobers := false
	enoughBags := true

	if enoughRobers || enoughBags {
		fmt.Println("Pegue tudo")
	} else {
		fmt.Println("Pegue o que Puder")
	}
	readyToGo := true

	if !readyToGo {
		fmt.Println("Ligue o carro")
	} else {
		fmt.Println("O que estamos esperando??")
	}

	amountStolen := 64650

	if amountStolen > 1000000 {
		fmt.Println("Nós acertamos na loteria")
	} else if amountStolen >= 5000 {
		fmt.Println("Pense em todos os doces que podemos comprar")
	} else {
		fmt.Println("Porque Fizemos isso")

	}

	switch name := "H. J. Simp"; name {
	case "butch":
		fmt.Println("Head to Robbers Roost.")
	case "Bonnie":
		fmt.Println("Stay put in Joplin")
	default:
		fmt.Println("Apenas se esconda")
	}
	//Randomizando e Semeando

	rand.Seed(time.Now().UnixNano())
	amountLeft := rand.Intn(10000)

	fmt.Println("amountLeft is:", amountLeft)

	if amountLeft > 5000 {
		fmt.Println("Em que devo gastar isso?")
	} else {
		fmt.Println("Cade meu dinheiro")
	}

}

package main

import "fmt"

func comendoTacos() {
	fmt.Println("Yum!")
}

func main() {
	comendoTacos()
}

// Alcance

func startGame() {
	instruções := "Pressione enter pra começar"
	fmt.Println(instruções)
}

func main() {
	startGame()
}

// Parâmetros

func calculoMarteAnos(terraAnos int) int {

	terraDias := terraAnos * 365
	MarteAnos := terraDias / 687
	return MarteAnos
}

func main() {
	minhaIdade := 25

	minhaIdadeMarte := calculoMarteAnos(minhaIdade)
	fmt.Println(minhaIdadeMarte)
}

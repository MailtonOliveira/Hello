package main

import (
	"fmt"
)

func main() {

	const Gravidade = 9.80

	var numOfFlavors int = 57

	var lancheFavorito string
	lancheFavorito = "Pão de Queijo"

	var description string
	description = "Meu Lanche favorito é " + lancheFavorito

	daysOnVacation := 5

	var hoursInDay = 24

	var cupsOfCoffeeConsumed int

	cupsOfCoffeeConsumed = 5

	CoolSneakers := 65.99
	niceNecklace := 82.98

	var taxCalculation float64

	animal := "Gato"
	animal1 := "Cachorro"

	template := "Eu gostaria de ter um %v."

	pet := "cachorrinho"

	var wish string

	wish = fmt.Sprintf(template, pet)

	preferência := fmt.Sprintln(animal, animal1)

	taxCalculation += CoolSneakers

	taxCalculation += niceNecklace

	taxCalculation *= .06000

	fmt.Println("Olá Kanela")

	fmt.Println(Gravidade)

	fmt.Println(description)

	fmt.Println(numOfFlavors)

	fmt.Println("Você Passou", daysOnVacation*hoursInDay, "horas de férias.")

	fmt.Println(cupsOfCoffeeConsumed)

	fmt.Printf("Meu Colar favorito custa $%.2f mas o Tênis está $%.2f.", niceNecklace, CoolSneakers)

	fmt.Println(" Total de $", CoolSneakers+niceNecklace, "com 6% de imposto sobre vendas", taxCalculation, "igual a", "$", CoolSneakers+niceNecklace+taxCalculation)

	fmt.Println("\n ")

	fmt.Print("Impressões ", "são ", "diferentes, ")

	fmt.Print("consegue ", "ver?")

	fmt.Println("\n ")

	fmt.Printf("Você é qual tipo de pessoa, %v ou %v? ", animal, animal1)

	fmt.Println("\n ")

	fmt.Printf("Trabalhando com um %T", animal)

	fmt.Println("\n ")

	fmt.Print(preferência)

	fmt.Println("\n ")

	fmt.Println(wish)

	fmt.Println("\n ")

}

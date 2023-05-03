package main

import "fmt"

func main() {
	fmt.Println("O que você gosta de fazer")

	var diverção string

	fmt.Scan(&diverção)

	fmt.Printf("Claro, Podemos %v %v %v, quando você quiser", diverção)

	fmt.Println("\n ")

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	caracoroa := rand.Intn(4)

	if caracoroa >= 3 {
		fmt.Println("Coroa")
	} else {
		fmt.Println("Cara")
	}
}

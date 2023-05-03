package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	eludedGuards := rand.Intn(100)
	isHeistOn := true

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		fmt.Println("Plan a better disguise next time?")
	}
	fmt.Println("isHeistOn is currently:", isHeistOn)

	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Println("Grab and GO!")
	} else if isHeistOn == false {
		fmt.Println("Vault can´t be opened")
	} else {
		fmt.Println("What's the combo to this lock again??")
	}
	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Print("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Print("Turns out vault doors don't open from the inside...")
		default:
			fmt.Print("Start the getaway car!")
		}
	}
	amtStolen := 10000 + rand.Intn(1000000)
	if isHeistOn {
		fmt.Println("$", amtStolen, "not bad")
	}
}

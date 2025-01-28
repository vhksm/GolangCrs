package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

func presentOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func endLineProgram() {
	fmt.Println("Goodbye!")
	fmt.Println("Thanks for choosing our bank")
}

func welcomeToBank() {
	fmt.Println("Welcome to Go Bank,", randomdata.Title(randomdata.Male), randomdata.FullName(randomdata.Male))
	fmt.Println("If you want to talk to a agent", randomdata.PhoneNumber())
	fmt.Println("This is your currency used:", randomdata.Currency())
}

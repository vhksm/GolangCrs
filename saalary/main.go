package main

import (
	"fmt"

	"errors"
)

func main() {

	var hour float64
	var days float64

	fmt.Println("Welcome to your hourly value calculator")
	fmt.Println("How many hours you work each day?")
	fmt.Scan(&hour)
	fmt.Println("How many days do you work in one month? ")
	fmt.Scan(&days)
	userSalary, err := getUserInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	salary := calculateHour(userSalary, hour, days)
	fmt.Println("This is your payment by hour: ", salary)

}

func getUserInfo() (float64, error) {
	var userSalary float64
	fmt.Print("Enter your salary: ")
	fmt.Scan(&userSalary)
	if userSalary <= 0 {
		return 0, errors.New("the Value must be greather than 0")
	}
	return userSalary, nil
}
func calculateHour(userSalary, hour, days float64) float64 {

	salary := userSalary / (hour * days)
	return salary
}

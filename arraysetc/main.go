package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Victor")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 2)

	courseRatings["golang"] = 7.5
	courseRatings["javascript"] = 5.0
	courseRatings["angular"] = 6.2

	courseRatings.output()

	//fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}

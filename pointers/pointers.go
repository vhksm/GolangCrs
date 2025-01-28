package main

import (
	"fmt"
)

func main() {
	age := 32 //regular variable

	/*var agePointer *int         //declarando a variavel especial do ponteiro
	agePointer = &age*/

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18
}

// usar pointer para evitar uma copia de um numero como o 32 desse codigo nao vale a pena, pois gasta muito tempo para criar o ponteiro e referenciar a ele, e usa-lo nas funcoes e esse numero quase nao ocupa espaco na memoria

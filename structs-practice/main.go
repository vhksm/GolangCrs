package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"vhksm.com/exerciciostructs/note"
	"vhksm.com/exerciciostructs/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// Display()
// }

type outputtable interface {
	saver
	Display()
}

// type outputtable interface {
// Save() error
// Display()
// }

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")

	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {
		return
	}
	outputData(userNote)
}

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
		return
	}
	stringVal, ok := value.(int)
	if ok {
		fmt.Println("String:", stringVal)
		return
	}
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	//default:
	//...
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving failed.")
		return err
	}
	switch data.(type) {
	case todo.Todo:
		fmt.Println("Saving the todo succeeded.")
	case note.Note:
		fmt.Println("Saving the note succeeded")
	default:
		fmt.Println("Saving succeeded")
	}
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	//nova funcao bufio newreader para ler o que foi escrito, usa o package os para especificar que vai ler diretamente da linha de comando (Stdin)=standard input
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n') // funcao reader ReadString indica quando o input vai parar de ser lido (\n) significa que quando quebrar a linha vai parar. usa a variavel text pois sera o texto que foi lido e vai ser retornado no final

	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n") // in windows a linebreak is created with \r\n
	text = strings.TrimSuffix(text, "\r")

	return text
}

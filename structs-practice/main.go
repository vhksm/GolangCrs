package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"vhksm.com/exerciciostructs/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.OutputUserDetails()

	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note suceeded.")

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

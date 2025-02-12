package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// (note Note) é um receiver, pois está attached ao struct Note, entao cria a variavel "note" de type Note (struct) para poder usar os valores que estão la dentro
func (note Note) Display() {
	fmt.Printf("\nYour note titled %v has the following content:\n\n%v\n\n", note.Title, note.Content)
}

// Não usa ponteiro para se referir a Note porque não edita a Note, se tivesse um metodo que edita teria de usar pointer "*" pois se nao usasse estaria editando a copia e nao o valor real.
// strings.ReplaceAll u.title troca o que foi digitado no titulo para o que foi escolhido (i.e. trocar " "(espaços vazios) por "_") strings.ToLower(fileName) troca todas as letras maiusculas que estavam no texto indicado pela variavel, no caso fileName para minusculas.
// encoding/json vai converter o conteudo no comando WriteFile para formato JSON, declara uma variavel com nome de json que ao usar o comando json.Marshal(u), sendo u a variavel para usar a struct Note, e a variavel json ja estara em formato de []byte
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644) /*0644 = Numero de permissao que possibilita o criador da pasta editar e ler tudo, mas outros users poderao apenas ler*/
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

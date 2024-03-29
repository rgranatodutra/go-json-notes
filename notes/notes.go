package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	CreatedAt     int64  `json:"created_at"`
	LastUpdatedAt int64  `json:"last_updated_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("can't create empty title or content")
	}

	return Note{
		Id:        uuid.NewString(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().Unix(),
	}, nil
}

func (n *Note) ChangeTitle(title string) {
	n.LastUpdatedAt = time.Now().Unix()
	n.Title = title
}

func (n *Note) ChangeContent(newContent string) {
	n.LastUpdatedAt = time.Now().Unix()
	n.Content = newContent
}

func (n *Note) Print() {
	createdAt := time.Unix(n.CreatedAt, 0).Format("02/01/2006 15:04:05")

	fmt.Printf("\nTitle             %s\n", n.Title)
	fmt.Printf("Created at        %s\n", createdAt)

	if n.LastUpdatedAt > 0 {
		lastUpdatedAt := time.Unix(n.LastUpdatedAt, 0).Format("02/01/2006 15:04:05")
		fmt.Printf("Last updated at   %s\n", lastUpdatedAt)
	}

	fmt.Printf("\n\"%s\"\n", n.Content)
}

func (n *Note) Save(savePath string) error {
	filename := strings.ReplaceAll(strings.ToLower(n.Title), " ", "_") + ".jn"
	filePath := filepath.Join(savePath, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	noteJSON, err := json.MarshalIndent(n, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling note to JSON: %w", err)
	}

	_, err = file.Write(noteJSON)
	if err != nil {
		return fmt.Errorf("error writing note to file: %w", err)
	}

	return nil
}

func Open(filePath string) (Note, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Note{}, fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	var note Note
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&note); err != nil {
		return Note{}, fmt.Errorf("erro ao decodificar o JSON: %w", err)
	}

	return note, nil
}

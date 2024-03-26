package notes

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Note struct {
	id            string
	title         string
	content       string
	createdAt     int64
	lastUpdatedAt int64
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("can't create empty title or content")
	}

	return Note{
		id:        uuid.NewString(),
		title:     title,
		content:   content,
		createdAt: time.Now().Unix(),
	}, nil
}

func (n *Note) ChangeTitle(title string) {
	n.lastUpdatedAt = time.Now().Unix()
	n.title = title
}

func (n *Note) ChangeContent(newContent string) {
	n.lastUpdatedAt = time.Now().Unix()
	n.content = newContent
}

func (n *Note) Print() {
	createdAt := time.Unix(n.createdAt, 0).Format("02/01/2006 15:04:05")

	fmt.Printf("\nTitle             %s\n", n.title)
	fmt.Printf("Created at        %s\n", createdAt)

	if n.lastUpdatedAt > 0 {
		lastUpdatedAt := time.Unix(n.lastUpdatedAt, 0).Format("02/01/2006 15:04:05")
		fmt.Printf("Last updated at   %s\n", lastUpdatedAt)
	}

	fmt.Printf("\n\"%s\"\n", n.content)
}

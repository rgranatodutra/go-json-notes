package notation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Notation struct {
	id            string
	title         string
	text          string
	createdAt     int64
	lastUpdatedAt int64
}

func New(title, text string) *Notation {
	return &Notation{
		id:        uuid.NewString(),
		title:     title,
		text:      text,
		createdAt: time.Now().Unix(),
	}
}

func (n *Notation) ChangeText(txt string) {
	n.lastUpdatedAt = time.Now().Unix()
	n.text = txt
}

func (n *Notation) ChangeTitle(title string) {
	n.lastUpdatedAt = time.Now().Unix()
	n.title = title
}

func (n *Notation) Print() {
	createdAt := time.Unix(n.createdAt, 0).Format("02/01/2006 15:04:05")

	fmt.Printf("\nTitle             %s\n", n.title)
	fmt.Printf("Created at        %s\n", createdAt)

	if n.lastUpdatedAt > 0 {
		lastUpdatedAt := time.Unix(n.lastUpdatedAt, 0).Format("02/01/2006 15:04:05")
		fmt.Printf("Last updated at   %s\n", lastUpdatedAt)
	}

	fmt.Printf("\n\"%s\"\n", n.text)
}

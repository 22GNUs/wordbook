package eduic

import "fmt"

// Explain repensent a word's explaination
type Explain struct {
	Word string
	Exp  string
}

func (e Explain) String() string {
	return fmt.Sprintf("%s: %s", e.Word, e.Exp)
}

type WordBody struct {
	Id       string   `json:"id"`
	Language string   `json:"language"`
	Words    []string `json:"words"`
}

type Response[D any] struct {
	Data    D      `json:"data"`
	Message string `json:"message"`
}

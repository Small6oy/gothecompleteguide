package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Missing value for text")
	}

	return Todo{
		text,
	}, nil
}

func (t Todo) Display() {
	fmt.Printf("Text: %v", t.Text)
}

func (n Todo) Save() error {

	fileName := strings.ReplaceAll(n.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

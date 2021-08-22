package models

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// ListItem represents a single item in the script filter.
type ListItem struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
}

// ScriptResponse represents a list of items that the script filter will provide to Alfred.
type ScriptResponse struct {
	Items  []ListItem `json:"items"`
	output io.Writer
}

type ScriptResponseOption func(*ScriptResponse) error

func WithOutput(writer io.Writer) ScriptResponseOption {
	return func(sr *ScriptResponse) error {
		sr.output = writer
		return nil
	}
}

func NewScriptResponse(opts ...ScriptResponseOption) (*ScriptResponse, error) {
	sr := &ScriptResponse{
		output: os.Stdout,
		Items:  make([]ListItem, 0),
	}

	for _, opt := range opts {
		if err := opt(sr); err != nil {
			return sr, err
		}
	}

	return sr, nil
}

func (sr *ScriptResponse) AddItem(item ListItem) {
	sr.Items = append(sr.Items, item)
}

func (sr *ScriptResponse) SendFeedback() {
	encoder := json.NewEncoder(sr.output)

	// the odds of an encoder faliure are tiny.
	if err := encoder.Encode(sr); err != nil {
		errResponse := &ScriptResponse{
			Items: []ListItem{
				{
					Title:    "Encoding Error!",
					Subtitle: err.Error(),
				},
			},
		}
		encoder.Encode(errResponse)
		log.Println(err)
	}
}

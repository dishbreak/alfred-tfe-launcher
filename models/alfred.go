package models

import (
	"encoding/json"
	"io"
	"os"
)

// ListItem represents a single item in the script filter.
type ListItem struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Valid    bool   `json:"valid"`
}

// ScriptResponse represents a list of items that the script filter will provide to Alfred.
type ScriptResponse struct {
	Items  []ListItem `json:"items"`
	output io.Writer
}

type ScriptResponseOption func(*ScriptResponse)

func WithOutput(writer io.Writer) ScriptResponseOption {
	return func(sr *ScriptResponse) {
		sr.output = writer
	}
}

func NewScriptResponse(opts ...ScriptResponseOption) *ScriptResponse {
	sr := &ScriptResponse{
		output: os.Stdout,
		Items:  make([]ListItem, 0),
	}

	for _, opt := range opts {
		opt(sr)
	}

	return sr
}

func (sr *ScriptResponse) AddItem(item ListItem) {
	sr.Items = append(sr.Items, item)
}

// SetError will write the error back to Alfred.
// Callers must return after calling this function!
func (sr *ScriptResponse) SetError(err error) {
	sr.Items = []ListItem{
		{
			Title:    "Encountered Error!",
			Subtitle: err.Error(),
			Valid:    false,
		},
	}
	sr.SendFeedback()
}

func (sr *ScriptResponse) SendFeedback() {
	encoder := json.NewEncoder(sr.output)
	encoder.Encode(sr)
}

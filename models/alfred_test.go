package models

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func badOption(sr *ScriptResponse) error {
	return errors.New("boooo!")
}

func TestNewScriptResponse(t *testing.T) {

	sr := NewScriptResponse()
	assert.Empty(t, sr.Items)
}

func TestScriptResponseAddItems(t *testing.T) {
	output := bytes.NewBufferString("")

	sr := NewScriptResponse(WithOutput(output))

	sr.AddItem(ListItem{
		Title:    "foo",
		Subtitle: "bar",
		Arg:      "fool",
	})

	sr.SendFeedback()
	assert.Equal(t, `{"items":[{"title":"foo","subtitle":"bar","arg":"fool","valid":false}]}
`, output.String())
}

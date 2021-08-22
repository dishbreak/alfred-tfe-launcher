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
	_, err := NewScriptResponse(badOption)
	assert.Error(t, err)

	sr, _ := NewScriptResponse()
	assert.Empty(t, sr.Items)
}

func TestScriptResponseAddItems(t *testing.T) {
	output := bytes.NewBufferString("")

	sr, err := NewScriptResponse(WithOutput(output))
	assert.Nil(t, err, "got error with script response")

	sr.AddItem(ListItem{
		Title:    "foo",
		Subtitle: "bar",
		Arg:      "fool",
	})

	sr.SendFeedback()
	assert.Equal(t, `{"items":[{"title":"foo","subtitle":"bar","arg":"fool"}]}
`, output.String())
}

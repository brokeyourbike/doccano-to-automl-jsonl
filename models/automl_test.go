package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCanMarshal(t *testing.T) {
	ann := Annotation{}
	ann.DisplayName = "h"
	ann.TextExtraction.TextSegment.StartOffset = 1
	ann.TextExtraction.TextSegment.EndOffset = 3

	a := AutoML{}
	a.TextSnippet.Content = "hello!"
	a.Annotations = []Annotation{ann}

	result, err := json.Marshal(&a)
	assert.NoError(t, err)

	j := `{"text_snippet":{"content":"hello!"},"annotations":[{"text_extraction":{"text_segment":{"start_offset":1,"end_offset":3}},"display_name":"h"}]}`
	assert.Equal(t, j, string(result))
}

func TestItCanEmmitAnnotations(t *testing.T) {
	a := AutoML{}
	a.TextSnippet.Content = "hello!"

	result, err := json.Marshal(&a)
	assert.NoError(t, err)

	j := `{"text_snippet":{"content":"hello!"}}`
	assert.Equal(t, j, string(result))
}

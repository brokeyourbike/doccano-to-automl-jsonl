package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItCanUnmarshal(t *testing.T) {
	j := `{"id": 1, "data": "Kharkiv. Example Steet", "label": [[9, 18, "Street"], [0, 7, "City"]]}`

	var d Doccano
	err := json.Unmarshal([]byte(j), &d)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), d.ID)
	assert.Equal(t, "Kharkiv. Example Steet", d.Data)
	assert.Equal(t, Label{Start: 9, End: 18, Name: "Street"}, d.Labels[0])
	assert.Equal(t, Label{Start: 0, End: 7, Name: "City"}, d.Labels[1])
}

func TestItCanConvert(t *testing.T) {
	d := Doccano{Data: "text", Labels: []Label{{Start: 1, End: 2, Name: "Letter"}}}

	a := d.Convert()

	assert.Equal(t, "text", a.TextSnippet.Content)
	assert.Equal(t, "Letter", a.Annotations[0].DisplayName)
	assert.Equal(t, int64(1), a.Annotations[0].TextExtraction.TextSegment.StartOffset)
	assert.Equal(t, int64(2), a.Annotations[0].TextExtraction.TextSegment.EndOffset)
}

package models

import "encoding/json"

type Label struct {
	Start int64
	End   int64
	Name  string
}

func (l *Label) UnmarshalJSON(p []byte) error {
	var tmp []json.RawMessage
	if err := json.Unmarshal(p, &tmp); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[0], &l.Start); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[1], &l.End); err != nil {
		return err
	}
	if err := json.Unmarshal(tmp[2], &l.Name); err != nil {
		return err
	}
	return nil
}

type Doccano struct {
	ID     int64   `json:"id"`
	Data   string  `json:"data"`
	Labels []Label `json:"label"`
}

func (d *Doccano) Convert() AutoML {
	a := AutoML{}
	a.TextSnippet.Content = d.Data
	a.Annotations = make([]Annotation, len(d.Labels))

	for i, l := range d.Labels {
		ann := Annotation{DisplayName: l.Name}
		ann.TextExtraction.TextSegment.StartOffset = l.Start
		ann.TextExtraction.TextSegment.EndOffset = l.End

		a.Annotations[i] = ann
	}

	return a
}

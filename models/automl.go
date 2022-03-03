package models

type Annotation struct {
	TextExtraction struct {
		TextSegment struct {
			StartOffset int64 `json:"start_offset"`
			EndOffset   int64 `json:"end_offset"`
		} `json:"text_segment"`
	} `json:"text_extraction"`
	DisplayName string `json:"display_name"`
}

type AutoML struct {
	TextSnippet struct {
		Content string `json:"content"`
	} `json:"text_snippet"`
	Annotations []Annotation `json:"annotations,omitempty"`
}

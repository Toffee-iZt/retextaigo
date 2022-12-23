package api

// Synonyms is synonyms response.
type Synonyms struct {
	Range   []int    `json:"range"`
	Changed []string `json:"changed"`
	Synonym []string `json:"synonym"`
}

// Paraphrased is paraphrase response.
type Paraphrased = []string

// Summarized is summarize response.
type Summarized = string

// Extended is extension response.
type Extended = []any

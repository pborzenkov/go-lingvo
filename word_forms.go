package lingvo

import (
	"context"
)

const (
	endpointWordForms = "api/v1/WordForms"
)

// Lexeme contains a single word with all its word forms.
type Lexeme struct {
	Lexeme       string   `json:"Lexem"`
	PartOfSpeech string   `json:"PartOfSpeech"`
	Paradigm     Paradigm `json:"ParadigmJson"`
}

// Paradigm contains word forms
type Paradigm struct {
	Name    string  `json:"Name"`
	Grammar string  `json:"Grammar"`
	Groups  []Group `json:"Groups"`
}

// Group contains a group of word forms
type Group struct {
	Name        string        `json:"Name"`
	Table       [][]TableCell `json:"Table"`
	ColumnCount int           `json:"ColumnCount"`
	RowCount    int           `json:"RowCount"`
}

// TableCell contains a single word form
type TableCell struct {
	Value  string `json:"Value"`
	Prefix string `json:"Prefix"`
	Row    string `json:"Row"`
}

// GetWordForms returns all word forms for word in language lang.
func (c *Client) GetWordForms(ctx context.Context, word string, lang Lang) ([]*Lexeme, error) {
	u, err := addOptions(endpointWordForms, option{"text", word}, option{"lang", lang.code()})
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	lexemes := new([]*Lexeme)
	err = c.Do(ctx, req, lexemes)
	if err != nil {
		return nil, err
	}

	return *lexemes, nil
}

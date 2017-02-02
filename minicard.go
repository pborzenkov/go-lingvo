package lingvo

import (
	"context"
)

const (
	endpointMinicard = "api/v1/Minicard"
)

// WordType is a type of word
type WordType int

// Types of words
const (
	None      WordType = 0
	ExactWord WordType = 1 << (iota - 1)
	LemmatizedVariant
	Subphrase
	SpellingVariant
)

// Minicard contains a minicard for a single word.
type Minicard struct {
	client *Client `json:-`

	SourceLanguage Lang     `json:"SourceLanguage"`
	TargetLanguage Lang     `json:"TargetLanguage"`
	Heading        string   `json:"Heading"`
	Translation    Word     `json:"Translation"`
	SeeAlso        []string `json:"SeeAlso"`
}

// Word contains a translation of a single word.
type Word struct {
	Heading      string   `json:"Heading"`
	Translation  string   `json:"Translation"`
	Dictionary   string   `json:"DictionaryName"`
	Sound        string   `json:"SoundName"`
	Type         WordType `json:"Type"`
	OriginalWord string   `json:"OriginalWord"`
}

// GetMinicard returns a minicard for word.
func (c *Client) GetMinicard(ctx context.Context, word string, from, to Lang) (*Minicard, error) {
	u, err := addOptions(endpointMinicard,
		option{"text", word},
		option{"srcLang", from.code()},
		option{"dstLang", to.code()},
	)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	minicard := new(Minicard)
	err = c.Do(ctx, req, minicard)
	if err != nil {
		return nil, err
	}
	minicard.client = c

	return minicard, nil
}

// GetSound returns a wav encoded sound for a word represented by the minicard.
func (m *Minicard) GetSound(ctx context.Context) ([]byte, error) {
	if m.Translation.Sound == "" {
		return nil, ErrNoSound
	}

	return m.client.GetSound(ctx, m.Translation.Dictionary, m.Translation.Sound)
}

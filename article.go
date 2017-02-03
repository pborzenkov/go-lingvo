package lingvo

//go:generate stringer -type=NodeType

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	endpointArticle = "api/v1/Article"
)

// NodeType is an article node type
type NodeType int

// Possible node types
const (
	Comment NodeType = iota
	Paragraph
	Text
	List
	ListItem
	Examples
	ExampleItem
	Example
	CardRefs
	CardRefItem
	CardRef
	Transcription
	Abbrev
	Caption
	Sound
	Ref
	Unsupported
)

var nodeTypes = []NodeType{
	Comment, Paragraph, Text, List, ListItem, Examples,
	ExampleItem, Example, CardRefs, CardRefItem, CardRef,
	Transcription, Abbrev, Caption, Sound, Ref, Unsupported,
}

var str2nodeType = make(map[string]NodeType)

func init() {
	for _, n := range nodeTypes {
		str2nodeType[n.String()] = n
	}
}

// MarshalJSON implement json.Marshaler interface
func (n NodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

// UnmarshalJSON implements json.Unmarshaler interface
func (n *NodeType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	nt, ok := str2nodeType[s]
	if !ok {
		return ErrInvalidNodeType
	}
	*n = nt
	return nil
}

// Article is a single article
type Article struct {
	Title      string         `json:"Title"`
	Markup     []*ArticleNode `json:"TitleMarkup"`
	Dictionary string         `json:"Dictionary"`
	ID         string         `json:"ArticleId"`
	Body       []*ArticleNode `json:"Body"`
}

// ArticleNode is a single entry in an article
type ArticleNode struct {
	Node       NodeType       `json:"Node"`
	Text       string         `json:"Text"`
	Dictionary string         `json:"Dictionary"`
	ID         string         `json:"ArticleId"`
	IsItalics  bool           `json:"IsItalics"`
	IsAccent   bool           `json:"IsAccent"`
	IsOptional bool           `json:"IsOptional"`
	Items      []*ArticleNode `json:"Items"`
	Markup     []*ArticleNode `json:"Markup"`
}

// GetArticle returns the article with heading from dict.
func (c *Client) GetArticle(ctx context.Context, heading, dict string, from, to Lang) (*Article, error) {
	dict = fmt.Sprintf("%s (%s-%s)", dict, from, to)
	u, err := addOptions(endpointArticle,
		option{"heading", heading},
		option{"dict", dict},
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

	article := new(Article)
	err = c.Do(ctx, req, article)
	if err != nil {
		return nil, err
	}

	return article, nil
}

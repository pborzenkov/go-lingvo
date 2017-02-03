package lingvo

import (
	"context"
	"strconv"
)

const (
	endpointTranslation = "api/v1/Translation"
)

// Translate returns all possible traslation of word from lang from to lang to.
func (c *Client) Translate(ctx context.Context, word string, from, to Lang, isCaseSensitive bool) ([]*Article, error) {
	u, err := addOptions(endpointTranslation,
		option{"text", word},
		option{"srcLang", from.code()},
		option{"dstLang", to.code()},
		option{"isCaseSensitive", strconv.FormatBool(isCaseSensitive)},
	)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	articles := new([]*Article)
	err = c.Do(ctx, req, articles)
	if err != nil {
		return nil, err
	}

	return *articles, nil
}

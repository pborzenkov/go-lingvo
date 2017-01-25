package lingvo

import (
	"golang.org/x/net/context"
)

const (
	endpointSuggests = "api/v1/Suggests"
)

// Suggest returns all spelling suggestions for word.
func (c *Client) Suggest(ctx context.Context, word string, from, to Lang) ([]string, error) {
	u, err := addOptions(endpointSuggests,
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

	suggestions := new([]string)
	err = c.Do(ctx, req, suggestions)
	if err != nil {
		return nil, err
	}

	return *suggestions, nil
}

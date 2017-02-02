package lingvo

import (
	"context"
	"encoding/base64"
)

const (
	endpointSound = "api/v1/Sound"
)

// GetSound returns a wav encoded sound file from dict
func (c *Client) GetSound(ctx context.Context, dict, file string) ([]byte, error) {
	u, err := addOptions(endpointSound,
		option{"dictionaryName", dict},
		option{"fileName", file},
	)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	var sound string
	err = c.Do(ctx, req, &sound)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(sound)
}

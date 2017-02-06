package lingvo

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

const (
	endpointSearch = "api/v1/Search"
)

// SearchZone defines where to perform the search
type SearchZone int

// Possible search zones
const (
	SearchHeading SearchZone = (1 << iota)
	SearchTranslation
	SearchExample
	SearchComment

	SearchAuto = (SearchHeading | SearchTranslation | SearchExample)
	SearchAll  = (SearchAuto | SearchComment)
)

// String implement Stringer interface
// Can't use 'stringer' here as it doesn't currently allow to strip prefix from
// the names.
func (z SearchZone) String() string {
	switch z {
	case SearchHeading:
		return "Heading"
	case SearchTranslation:
		return "Translation"
	case SearchExample:
		return "Example"
	case SearchComment:
		return "Comment"
	case SearchAuto:
		return "Auto"
	case SearchAll:
		return "All"
	default:
		return fmt.Sprintf("SearchZone(%d)", z)
	}
}

// code returns strings representation of search zone code
func (z SearchZone) code() string {
	return strconv.Itoa(int(z))
}

var searchZones = []SearchZone{
	SearchHeading, SearchTranslation, SearchExample, SearchComment,
	SearchAuto, SearchAll,
}

var str2searchZone = make(map[string]SearchZone)

func init() {
	for _, z := range searchZones {
		str2searchZone[strings.ToLower(z.String())] = z
	}
}

// Set implements flag.Value interface
func (z *SearchZone) Set(str string) error {
	val, ok := str2searchZone[str]
	if !ok {
		return ErrInvalidSearchZone
	}
	*z = val
	return nil
}

// SearchResult is the result of a search
type SearchResult struct {
	Items       []*Article `json:"Items"`
	Total       int        `json:"TotalCount"`
	HasNextPage bool       `json:"HasNextPage"`
}

// Search performs the search
func (c *Client) Search(ctx context.Context, word string, from, to Lang, zone SearchZone, start, pageSize int) (*SearchResult, error) {
	u, err := addOptions(endpointSearch,
		option{"text", word},
		option{"srcLang", from.code()},
		option{"dstLang", to.code()},
		option{"searchZone", zone.code()},
		option{"startIndex", strconv.Itoa(start)},
		option{"pageSize", strconv.Itoa(pageSize)},
	)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	result := new(SearchResult)
	err = c.Do(ctx, req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

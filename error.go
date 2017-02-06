package lingvo

import (
	"errors"
)

// Possible error values
var (
	ErrUnsupportedLanguage = errors.New("lingvo: unsupported language")
	ErrNoSound             = errors.New("lingvo: no sound")
	ErrInvalidNodeType     = errors.New("lingvo: invalid node type")
	ErrInvalidSearchZone   = errors.New("lingvo: invalid search zone")
)

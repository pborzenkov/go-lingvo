package lingvo

import (
	"errors"
)

// Possible error values
var (
	ErrUnsupportedLanguage = errors.New("lingvo: unsupported language")
	ErrNoSound             = errors.New("lingvo: no sound")
)
//go:generate stringer -type=Lang

package lingvo

import (
	"strconv"
	"strings"
)

// Lang is a language to translate to/from.
type Lang int

// Language codes for all supported languages.
const (
	Ch Lang = 1028
	Da Lang = 1030
	De Lang = 1031
	El Lang = 1032
	En Lang = 1033
	Es Lang = 1034
	Fr Lang = 1036
	It Lang = 1040
	Pl Lang = 1045
	Ru Lang = 1049
	Uk Lang = 1058
	Kk Lang = 1087
	Tt Lang = 1092
	La Lang = 1142
)

// code returns string representation of l language code
func (l Lang) code() string {
	return strconv.Itoa(int(l))
}

// Languages is a list of all supported languages.
var Languages = []Lang{Ch, Da, De, El, En, Es, Fr, It, Pl, Ru, Uk, Kk, Tt, La}

var abbr2lang = make(map[string]Lang)

func init() {
	for _, l := range Languages {
		abbr2lang[strings.ToLower(l.String())] = l
	}
}

// LangByAbbr returns language identified by abbr abbreviation.
func LangByAbbr(abbr string) (Lang, error) {
	l, ok := abbr2lang[strings.ToLower(abbr)]
	if !ok {
		return 0, ErrUnsupportedLanguage
	}
	return l, nil
}

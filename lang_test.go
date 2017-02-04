package lingvo

import (
	"testing"
)

func TestLangByAbbr(t *testing.T) {
	for _, l := range Languages {
		abbr := l.String()
		rl, err := LangByAbbrev(abbr)
		if err != nil {
			t.Errorf("unexpected error '%v' for lang '%s'", err, l)
		}
		if l != rl {
			t.Errorf("unexpected language for abbr '%s', want = %s, got = %s", abbr, l, rl)
		}
	}

	_, err := LangByAbbrev("??")
	if err != ErrUnsupportedLanguage {
		t.Errorf("unexpected language for '??', want = %v, got = %v", ErrUnsupportedLanguage, err)
	}
}

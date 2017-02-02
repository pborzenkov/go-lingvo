package lingvo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetMinicard(t *testing.T) {
	setup()
	defer teardown()

	// To make reflect.DeepEqual happy
	testMinicard.client = client

	mux.HandleFunc("/"+endpointMinicard, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"text":    "冬",
			"srcLang": Ch.code(),
			"dstLang": Ru.code(),
		})

		fmt.Fprintf(w, testMinicardJSON)
	})

	s, err := client.GetMinicard(context.Background(), "冬", Ch, Ru)
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if got, want := s, testMinicard; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected result, want = %+v, got = %+v", want, got)
	}
}

var testMinicard = &Minicard{
	SourceLanguage: Ch,
	TargetLanguage: Ru,
	Heading:        "冬",
	Translation: Word{
		Heading:      "冬",
		Translation:  "зима́; зи́мний",
		Dictionary:   "Universal (Ch-Ru)",
		Sound:        "",
		Type:         ExactWord,
		OriginalWord: "",
	},
	SeeAlso: []string{},
}

var testMinicardJSON = `
{
  "SourceLanguage": 1028,
  "TargetLanguage": 1049,
  "Heading": "冬",
  "Translation": {
    "Heading": "冬",
    "Translation": "зима́; зи́мний",
    "DictionaryName": "Universal (Ch-Ru)",
    "SoundName": "",
    "Type": 1,
    "OriginalWord": ""
  },
  "SeeAlso": []
}
`

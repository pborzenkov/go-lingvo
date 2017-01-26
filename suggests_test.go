package lingvo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSuggest(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+endpointSuggests, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"text":    "helo",
			"srcLang": En.code(),
			"dstLang": Ru.code(),
		})

		fmt.Fprintf(w, testSuggestsJSON)
	})

	s, err := client.Suggest(context.Background(), "helo", En, Ru)
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if got, want := s, testSuggests; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected result, want = %+v, got = %+v", want, got)
	}
}

var testSuggests = []string{
	"below",
	"cel",
	"cello",
	"CEO",
	"chela",
	"del",
	"Deo",
	"felon",
	"filo",
	"gel",
	"geo",
	"hale",
	"hall",
	"hallo",
	"halm",
	"halo",
	"halt",
	"he",
	"heal",
	"heel",
	"heir",
	"Hekla",
	"held",
	"Helen",
	"helix",
	"hell",
	"hello",
	"helm",
	"helot",
	"help",
	"helve",
	"heme",
	"here",
	"hero",
	"Herod",
	"heron",
	"hila",
	"hili",
	"hill",
	"hilt",
	"hl",
	"ho",
	"hobo",
	"hole",
	"holer",
	"holey",
	"hollo",
	"homo",
	"hula",
	"hullo",
	"hypo",
	"jello",
	"kilo",
	"melon",
	"Milo",
	"pel",
	"silo",
	"tel",
	"telco",
	"Teo",
}

var testSuggestsJSON = `
[
  "below",
  "cel",
  "cello",
  "CEO",
  "chela",
  "del",
  "Deo",
  "felon",
  "filo",
  "gel",
  "geo",
  "hale",
  "hall",
  "hallo",
  "halm",
  "halo",
  "halt",
  "he",
  "heal",
  "heel",
  "heir",
  "Hekla",
  "held",
  "Helen",
  "helix",
  "hell",
  "hello",
  "helm",
  "helot",
  "help",
  "helve",
  "heme",
  "here",
  "hero",
  "Herod",
  "heron",
  "hila",
  "hili",
  "hill",
  "hilt",
  "hl",
  "ho",
  "hobo",
  "hole",
  "holer",
  "holey",
  "hollo",
  "homo",
  "hula",
  "hullo",
  "hypo",
  "jello",
  "kilo",
  "melon",
  "Milo",
  "pel",
  "silo",
  "tel",
  "telco",
  "Teo"
]
`

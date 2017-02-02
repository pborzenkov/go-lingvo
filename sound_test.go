package lingvo

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var (
	testSoundJSON string
	testSound     []byte
)

func initFixtures(t *testing.T) {
	js, err := ioutil.ReadFile("fixtures/bang.json")
	if err != nil {
		t.Fatalf("failed to read \"fixtures/bang.json\": %v", err)
	}
	testSoundJSON = string(js)

	testSound, err = ioutil.ReadFile("fixtures/bang.wav")
	if err != nil {
		t.Fatalf("failed to read \"fixtures/bang.wav\": %v", err)
	}
}

func TestGetSound(t *testing.T) {
	setup()
	defer teardown()

	initFixtures(t)

	mux.HandleFunc("/"+endpointSound, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"dictionaryName": "LingvoUniversal (En-Ru)",
			"fileName":       "bang.wav",
		})

		fmt.Fprintf(w, testSoundJSON)
	})

	s, err := client.GetSound(context.Background(), "LingvoUniversal (En-Ru)", "bang.wav")
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if got, want := s, testSound; !bytes.Equal(got, want) {
		t.Errorf("unexpected result, want = %x, got = %x", want, got)
	}
}

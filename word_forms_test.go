package lingvo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestGetWordForms(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+endpointWordForms, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"text": "колено",
			"lang": Ru.code(),
		})

		fmt.Fprintf(w, testWordFormsJSON)
	})

	wf, err := client.GetWordForms(context.Background(), "колено", Ru)
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if got, want := wf, testWordForms; !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected result, want = %+v, got = %+v", want, got)
	}
}

var testWordForms = []*Lexeme{
	&Lexeme{
		Lexeme:       "колено",
		PartOfSpeech: "существительное",
		Paradigm: Paradigm{
			Name:    "колено",
			Grammar: "существительное, неодушевлённое, средний род",
			Groups: []Group{
				Group{
					Name: "",
					Table: [][]TableCell{
						[]TableCell{
							TableCell{Value: "", Prefix: "", Row: ""},
							TableCell{Value: "Ед. ч.", Prefix: "", Row: ""},
							TableCell{Value: "Мн. ч.", Prefix: "", Row: ""},
						},
						[]TableCell{
							TableCell{Value: "Именительный", Prefix: "", Row: ""},
							TableCell{Value: "колено", Prefix: "", Row: ""},
							TableCell{Value: "колена", Prefix: "", Row: ""},
						},
						[]TableCell{
							TableCell{Value: "Родительный", Prefix: "", Row: ""},
							TableCell{Value: "колена", Prefix: "", Row: ""},
							TableCell{Value: "колен", Prefix: "", Row: ""},
						},
						[]TableCell{
							TableCell{Value: "Дательный", Prefix: "", Row: ""},
							TableCell{Value: "колену", Prefix: "", Row: ""},
							TableCell{Value: "коленам", Prefix: "", Row: ""},
						},
						[]TableCell{
							TableCell{Value: "Винительный", Prefix: "", Row: ""},
							TableCell{Value: "колено", Prefix: "", Row: ""},
							TableCell{Value: "колена", Prefix: "", Row: ""},
						},
						[]TableCell{
							TableCell{Value: "Творительный", Prefix: "", Row: ""},
							TableCell{Value: "коленом", Prefix: "", Row: ""},
							TableCell{Value: "коленами", Prefix: "", Row: ""},
						},
						[]TableCell{
							TableCell{Value: "Предложный", Prefix: "", Row: ""},
							TableCell{Value: "колене", Prefix: "", Row: ""},
							TableCell{Value: "коленах", Prefix: "", Row: ""},
						},
					},
					ColumnCount: 3,
					RowCount:    7,
				},
			},
		},
	},
}

var testWordFormsJSON = `
[
  {
    "Lexem": "колено",
    "PartOfSpeech": "существительное",
    "ParadigmJson": {
      "Name": "колено",
      "Grammar": "существительное, неодушевлённое, средний род",
      "Groups": [
        {
          "Name": "",
          "Table": [
            [
              {
                "Value": "",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "Ед. ч.",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "Мн. ч.",
                "Prefix": null,
                "Row": null
              }
            ],
            [
              {
                "Value": "Именительный",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "колено",
                "Prefix": "",
                "Row": null
              },
              {
                "Value": "колена",
                "Prefix": "",
                "Row": null
              }
            ],
            [
              {
                "Value": "Родительный",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "колена",
                "Prefix": "",
                "Row": null
              },
              {
                "Value": "колен",
                "Prefix": "",
                "Row": null
              }
            ],
            [
              {
                "Value": "Дательный",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "колену",
                "Prefix": "",
                "Row": null
              },
              {
                "Value": "коленам",
                "Prefix": "",
                "Row": null
              }
            ],
            [
              {
                "Value": "Винительный",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "колено",
                "Prefix": "",
                "Row": null
              },
              {
                "Value": "колена",
                "Prefix": "",
                "Row": null
              }
            ],
            [
              {
                "Value": "Творительный",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "коленом",
                "Prefix": "",
                "Row": null
              },
              {
                "Value": "коленами",
                "Prefix": "",
                "Row": null
              }
            ],
            [
              {
                "Value": "Предложный",
                "Prefix": null,
                "Row": null
              },
              {
                "Value": "колене",
                "Prefix": "",
                "Row": null
              },
              {
                "Value": "коленах",
                "Prefix": "",
                "Row": null
              }
            ]
          ],
          "ColumnCount": 3,
          "RowCount": 7
        }
      ]
    }
  }
]`

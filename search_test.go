package lingvo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestSearch(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+endpointSearch, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"text":       "board",
			"srcLang":    En.code(),
			"dstLang":    Ru.code(),
			"searchZone": SearchAll.code(),
			"startIndex": "0",
			"pageSize":   "1",
		})

		fmt.Fprintf(w, testSearchJSON)
	})

	s, err := client.Search(context.Background(), "board", En, Ru, SearchAll, 0, 1)
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if got, want := s, testSearch; !reflect.DeepEqual(got, want) {
		diff := pretty.Compare(got, want)
		t.Errorf("unexpected result, diff = %s", diff)
	}
}

var testSearch = &SearchResult{
	Items: []*Article{
		&Article{
			Title: "Board",
			Markup: []*ArticleNode{
				&ArticleNode{
					IsAccent:   false,
					IsItalics:  false,
					IsOptional: false,
					Node:       Text,
					Text:       "Board",
				},
			},
			Dictionary: "Engineering (En-Ru)",
			ID:         "Engineering (En-Ru)__Board",
			Body: []*ArticleNode{
				&ArticleNode{
					IsOptional: false,
					Markup: []*ArticleNode{
						&ArticleNode{
							FullText:   "употребляется в сочетании",
							IsOptional: false,
							Node:       Abbrev,
							Text:       "в соч.",
						},
					},
					Node: Paragraph,
				},
				&ArticleNode{
					IsOptional: false,
					Items: []*ArticleNode{
						&ArticleNode{
							IsOptional: false,
							Markup: []*ArticleNode{
								&ArticleNode{
									ID:         "Engineering (En-Ru)__Board of Standards Review",
									Dictionary: "Engineering (En-Ru)",
									IsOptional: false,
									Node:       CardRef,
									Text:       "Board of Standards Review",
								},
							},
							Node: CardRefItem,
						},
						&ArticleNode{
							IsOptional: false,
							Markup: []*ArticleNode{
								&ArticleNode{
									ID:         "Engineering (En-Ru)__ISO Technical Board",
									Dictionary: "Engineering (En-Ru)",
									IsOptional: false,
									Node:       CardRef,
									Text:       "ISO Technical Board",
								},
							},
							Node: CardRefItem,
						},
					},
					Node: CardRefs,
				},
			},
		},
	},
	Total:       1223,
	HasNextPage: true,
}

var testSearchJSON = `
{
  "HasNextPage": true,
  "Items": [
    {
      "ArticleId": "Engineering (En-Ru)__Board",
      "Body": [
        {
          "IsOptional": false,
          "Markup": [
            {
              "FullText": "употребляется в сочетании",
              "IsOptional": false,
              "Node": "Abbrev",
              "Text": "в соч."
            }
          ],
          "Node": "Paragraph",
          "Text": null
        },
        {
          "IsOptional": false,
          "Items": [
            {
              "IsOptional": false,
              "Markup": [
                {
                  "ArticleId": "Engineering (En-Ru)__Board of Standards Review",
                  "Dictionary": "Engineering (En-Ru)",
                  "IsOptional": false,
                  "Node": "CardRef",
                  "Text": "Board of Standards Review"
                }
              ],
              "Node": "CardRefItem",
              "Text": null
            },
            {
              "IsOptional": false,
              "Markup": [
                {
                  "ArticleId": "Engineering (En-Ru)__ISO Technical Board",
                  "Dictionary": "Engineering (En-Ru)",
                  "IsOptional": false,
                  "Node": "CardRef",
                  "Text": "ISO Technical Board"
                }
              ],
              "Node": "CardRefItem",
              "Text": null
            }
          ],
          "Node": "CardRefs",
          "Text": null,
          "Type": null
        }
      ],
      "Dictionary": "Engineering (En-Ru)",
      "Title": "Board",
      "TitleMarkup": [
        {
          "IsAccent": false,
          "IsItalics": false,
          "IsOptional": false,
          "Node": "Text",
          "Text": "Board"
        }
      ]
    }
  ],
  "TotalCount": 1223
}
`

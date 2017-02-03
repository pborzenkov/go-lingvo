package lingvo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestGetArticle(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+endpointArticle, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"heading": "pin",
			"dict":    "Electronics (En-Ru)",
			"srcLang": En.code(),
			"dstLang": Ru.code(),
		})

		fmt.Fprintf(w, testArticleJSON)
	})

	s, err := client.GetArticle(context.Background(), "pin", "Electronics", En, Ru)
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if got, want := s, testArticle; !reflect.DeepEqual(got, want) {
		diff := pretty.Compare(got, want)
		t.Errorf("unexpected result, diff = %s", diff)
	}
}

var testArticle = &Article{
	Title: "pin",
	Markup: []*ArticleNode{
		&ArticleNode{
			Node:       Text,
			Text:       "pin",
			IsItalics:  false,
			IsAccent:   false,
			IsOptional: false,
		},
	},
	Dictionary: "Electronics (En-Ru)",
	ID:         "Electronics (En-Ru)__pin",
	Body: []*ArticleNode{
		&ArticleNode{
			Items: []*ArticleNode{
				&ArticleNode{
					Markup: []*ArticleNode{
						&ArticleNode{
							Markup: []*ArticleNode{
								&ArticleNode{
									IsItalics:  false,
									IsAccent:   false,
									Node:       Text,
									Text:       "штырь, штырёк ",
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											IsItalics:  false,
											IsAccent:   false,
											Node:       Text,
											Text:       "(",
											IsOptional: false,
										},
										&ArticleNode{
											IsItalics:  true,
											IsAccent:   false,
											Node:       Text,
											Text:       "напр. электрического соединителя",
											IsOptional: false,
										},
										&ArticleNode{
											IsItalics:  false,
											IsAccent:   false,
											Node:       Text,
											Text:       ")",
											IsOptional: false,
										},
									},
									Node:       Comment,
									IsOptional: false,
								},
							},
							Node:       Paragraph,
							IsOptional: false,
						},
					},
					Node:       ListItem,
					IsOptional: false,
				},
				&ArticleNode{
					Markup: []*ArticleNode{
						&ArticleNode{
							Markup: []*ArticleNode{
								&ArticleNode{
									IsItalics:  false,
									IsAccent:   false,
									Node:       Text,
									Text:       "болт; винт",
									IsOptional: false,
								},
							},
							Node:       Paragraph,
							IsOptional: false,
						},
					},
					Node:       ListItem,
					IsOptional: false,
				},
				&ArticleNode{
					Markup: []*ArticleNode{
						&ArticleNode{
							Markup: []*ArticleNode{
								&ArticleNode{
									IsItalics:  false,
									IsAccent:   false,
									Node:       Text,
									Text:       "шпонка; шплинт; штифт; палец",
									IsOptional: false,
								},
							},
							Node:       Paragraph,
							IsOptional: false,
						},
					},
					Node:       ListItem,
					IsOptional: false,
				},
				&ArticleNode{
					Markup: []*ArticleNode{
						&ArticleNode{
							Markup: []*ArticleNode{
								&ArticleNode{
									IsItalics:  false,
									IsAccent:   false,
									Node:       Text,
									Text:       "игла ",
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											IsItalics:  false,
											IsAccent:   false,
											Node:       Text,
											Text:       "(",
											IsOptional: false,
										},
										&ArticleNode{
											IsItalics:  true,
											IsAccent:   false,
											Node:       Text,
											Text:       "напр. матричного принтера",
											IsOptional: false,
										},
										&ArticleNode{
											IsItalics:  false,
											IsAccent:   false,
											Node:       Text,
											Text:       ")",
											IsOptional: false,
										},
									},
									Node:       Comment,
									IsOptional: false,
								},
							},
							Node:       Paragraph,
							IsOptional: false,
						},
						&ArticleNode{
							Items: []*ArticleNode{
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__alignment pin",
											Node:       CardRef,
											Text:       "alignment pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__base pin",
											Node:       CardRef,
											Text:       "base pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__control pin",
											Node:       CardRef,
											Text:       "control pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__cylindrical pin",
											Node:       CardRef,
											Text:       "cylindrical pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__disk cartridge lock pins",
											Node:       CardRef,
											Text:       "disk cartridge lock pins",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__drive pin",
											Node:       CardRef,
											Text:       "drive pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__grooved pin",
											Node:       CardRef,
											Text:       "grooved pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__guide pin",
											Node:       CardRef,
											Text:       "guide pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__input pin",
											Node:       CardRef,
											Text:       "input pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__input/output pins",
											Node:       CardRef,
											Text:       "input/output pins",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__matching pin",
											Node:       CardRef,
											Text:       "matching pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__outputpin",
											Node:       CardRef,
											Text:       "outputpin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__split pin",
											Node:       CardRef,
											Text:       "split pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__tapered pin",
											Node:       CardRef,
											Text:       "tapered pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
								&ArticleNode{
									Markup: []*ArticleNode{
										&ArticleNode{
											Dictionary: "Electronics (En-Ru)",
											ID:         "Electronics (En-Ru)__tube pin",
											Node:       CardRef,
											Text:       "tube pin",
											IsOptional: false,
										},
									},
									Node:       CardRefItem,
									IsOptional: false,
								},
							},
							Node:       CardRefs,
							IsOptional: true,
						},
					},
					Node:       ListItem,
					IsOptional: false,
				},
			},
			Node:       List,
			IsOptional: false,
		},
	},
}

var testArticleJSON = `
{
  "Title": "pin",
  "TitleMarkup": [
    {
      "IsItalics": false,
      "IsAccent": false,
      "Node": "Text",
      "Text": "pin",
      "IsOptional": false
    }
  ],
  "Dictionary": "Electronics (En-Ru)",
  "ArticleId": "Electronics (En-Ru)__pin",
  "Body": [
    {
      "Type": 3,
      "Items": [
        {
          "Markup": [
            {
              "Markup": [
                {
                  "IsItalics": false,
                  "IsAccent": false,
                  "Node": "Text",
                  "Text": "штырь, штырёк ",
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "IsItalics": false,
                      "IsAccent": false,
                      "Node": "Text",
                      "Text": "(",
                      "IsOptional": false
                    },
                    {
                      "IsItalics": true,
                      "IsAccent": false,
                      "Node": "Text",
                      "Text": "напр. электрического соединителя",
                      "IsOptional": false
                    },
                    {
                      "IsItalics": false,
                      "IsAccent": false,
                      "Node": "Text",
                      "Text": ")",
                      "IsOptional": false
                    }
                  ],
                  "Node": "Comment",
                  "Text": null,
                  "IsOptional": false
                }
              ],
              "Node": "Paragraph",
              "Text": null,
              "IsOptional": false
            }
          ],
          "Node": "ListItem",
          "Text": null,
          "IsOptional": false
        },
        {
          "Markup": [
            {
              "Markup": [
                {
                  "IsItalics": false,
                  "IsAccent": false,
                  "Node": "Text",
                  "Text": "болт; винт",
                  "IsOptional": false
                }
              ],
              "Node": "Paragraph",
              "Text": null,
              "IsOptional": false
            }
          ],
          "Node": "ListItem",
          "Text": null,
          "IsOptional": false
        },
        {
          "Markup": [
            {
              "Markup": [
                {
                  "IsItalics": false,
                  "IsAccent": false,
                  "Node": "Text",
                  "Text": "шпонка; шплинт; штифт; палец",
                  "IsOptional": false
                }
              ],
              "Node": "Paragraph",
              "Text": null,
              "IsOptional": false
            }
          ],
          "Node": "ListItem",
          "Text": null,
          "IsOptional": false
        },
        {
          "Markup": [
            {
              "Markup": [
                {
                  "IsItalics": false,
                  "IsAccent": false,
                  "Node": "Text",
                  "Text": "игла ",
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "IsItalics": false,
                      "IsAccent": false,
                      "Node": "Text",
                      "Text": "(",
                      "IsOptional": false
                    },
                    {
                      "IsItalics": true,
                      "IsAccent": false,
                      "Node": "Text",
                      "Text": "напр. матричного принтера",
                      "IsOptional": false
                    },
                    {
                      "IsItalics": false,
                      "IsAccent": false,
                      "Node": "Text",
                      "Text": ")",
                      "IsOptional": false
                    }
                  ],
                  "Node": "Comment",
                  "Text": null,
                  "IsOptional": false
                }
              ],
              "Node": "Paragraph",
              "Text": null,
              "IsOptional": false
            },
            {
              "Type": null,
              "Items": [
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__alignment pin",
                      "Node": "CardRef",
                      "Text": "alignment pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__base pin",
                      "Node": "CardRef",
                      "Text": "base pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__control pin",
                      "Node": "CardRef",
                      "Text": "control pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__cylindrical pin",
                      "Node": "CardRef",
                      "Text": "cylindrical pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__disk cartridge lock pins",
                      "Node": "CardRef",
                      "Text": "disk cartridge lock pins",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__drive pin",
                      "Node": "CardRef",
                      "Text": "drive pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__grooved pin",
                      "Node": "CardRef",
                      "Text": "grooved pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__guide pin",
                      "Node": "CardRef",
                      "Text": "guide pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__input pin",
                      "Node": "CardRef",
                      "Text": "input pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__input/output pins",
                      "Node": "CardRef",
                      "Text": "input/output pins",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__matching pin",
                      "Node": "CardRef",
                      "Text": "matching pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__outputpin",
                      "Node": "CardRef",
                      "Text": "outputpin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__split pin",
                      "Node": "CardRef",
                      "Text": "split pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__tapered pin",
                      "Node": "CardRef",
                      "Text": "tapered pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                },
                {
                  "Markup": [
                    {
                      "Dictionary": "Electronics (En-Ru)",
                      "ArticleId": "Electronics (En-Ru)__tube pin",
                      "Node": "CardRef",
                      "Text": "tube pin",
                      "IsOptional": false
                    }
                  ],
                  "Node": "CardRefItem",
                  "Text": null,
                  "IsOptional": false
                }
              ],
              "Node": "CardRefs",
              "Text": null,
              "IsOptional": true
            }
          ],
          "Node": "ListItem",
          "Text": null,
          "IsOptional": false
        }
      ],
      "Node": "List",
      "Text": null,
      "IsOptional": false
    }
  ]
}
`

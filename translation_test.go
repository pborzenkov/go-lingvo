package lingvo

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestTranslate(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+endpointTranslation, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"text":            "plum",
			"srcLang":         En.code(),
			"dstLang":         Ru.code(),
			"isCaseSensitive": "false",
		})

		fmt.Fprintf(w, testTranslationJSON)
	})

	s, err := client.Translate(context.Background(), "plum", En, Ru, false)
	if err != nil {
		t.Errorf("unexpected error '%v'", err)
	}

	if got, want := s, testTranslation; !reflect.DeepEqual(got, want) {
		diff := pretty.Compare(got, want)

		t.Errorf("unexpected result, diff = %s", diff)
	}
}

var testTranslation = []*Article{
	&Article{
		Title: "plum",
		Markup: []*ArticleNode{
			&ArticleNode{
				Node:       Text,
				Text:       "plum",
				IsItalics:  false,
				IsAccent:   false,
				IsOptional: false,
			},
		},
		Dictionary: "LingvoUniversal (En-Ru)",
		ID:         "LingvoUniversal (En-Ru)__plum",
		Body: []*ArticleNode{
			&ArticleNode{
				Markup: []*ArticleNode{
					&ArticleNode{
						Node:       Transcription,
						Text:       "plʌm",
						IsOptional: false,
					},
					&ArticleNode{
						IsItalics:  false,
						IsAccent:   false,
						Node:       Text,
						Text:       " ",
						IsOptional: false,
					},
					&ArticleNode{
						FullText:   "британский вариант английского языка; употребляется в Великобритании",
						Node:       Abbrev,
						Text:       "брит.",
						IsOptional: false,
					},
					&ArticleNode{
						IsItalics:  false,
						IsAccent:   false,
						Node:       Text,
						Text:       " ",
						IsOptional: false,
					},
					&ArticleNode{
						FileName:   "plum.wav",
						Node:       Sound,
						IsOptional: false,
					},
					&ArticleNode{
						IsItalics:  false,
						IsAccent:   false,
						Node:       Text,
						Text:       " / ",
						IsOptional: false,
					},
					&ArticleNode{
						FullText:   "американский вариант английского языка; употребляется в США",
						Node:       Abbrev,
						Text:       "амер.",
						IsOptional: false,
					},
					&ArticleNode{
						IsItalics:  false,
						IsAccent:   false,
						Node:       Text,
						Text:       " ",
						IsOptional: false,
					},
					&ArticleNode{
						FileName:   "9612.wav",
						Node:       Sound,
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
								Markup: []*ArticleNode{
									&ArticleNode{
										FullText:   "имя существительное",
										Node:       Abbrev,
										Text:       "сущ.",
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
												Items: []*ArticleNode{
													&ArticleNode{
														Markup: []*ArticleNode{
															&ArticleNode{
																Markup: []*ArticleNode{
																	&ArticleNode{
																		FullText:   "ботаника",
																		Node:       Abbrev,
																		Text:       "бот.",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "; = plum tree ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "слива ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		Markup: []*ArticleNode{
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "(",
																				IsOptional: false,
																			},
																			&ArticleNode{
																				FullText:   "латинский язык",
																				Node:       Abbrev,
																				Text:       "лат.",
																				IsOptional: false,
																			},
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "Prunus; ",
																				IsOptional: false,
																			},
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "вид плодового дерева)",
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
																		Text:       "слива ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		Markup: []*ArticleNode{
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "(плод)",
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
																				Dictionary: "LingvoUniversal (En-Ru)",
																				ID:         "LingvoUniversal (En-Ru)__French plum",
																				Node:       CardRef,
																				Text:       "French plum",
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
										Node:       ListItem,
										IsOptional: false,
									},
									&ArticleNode{
										Markup: []*ArticleNode{
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
																		Text:       "изюм",
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
																		Text:       "сорт десертного винограда ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		Markup: []*ArticleNode{
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "(предназначенный для изготовления изюма)",
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
																		FullText:   "американский вариант английского языка; употребляется в США",
																		Node:       Abbrev,
																		Text:       "амер.",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "круглый леденец ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		Markup: []*ArticleNode{
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "(по форме напоминающий сливу)",
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
																Node:       Caption,
																Text:       "Syn:",
																IsOptional: true,
															},
															&ArticleNode{
																Markup: []*ArticleNode{
																	&ArticleNode{
																		Dictionary: "LingvoUniversal (En-Ru)",
																		ID:         "LingvoUniversal (En-Ru)__sugarplum",
																		Node:       CardRef,
																		Text:       "sugarplum",
																		IsOptional: false,
																	},
																},
																Node:       Paragraph,
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
														Text:       "лакомый кусочек; нечто самое лучшее; сливки",
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
																Markup: []*ArticleNode{
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "to pick / take the plum ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "— снимать сливки, отбирать самое лучшее",
																		IsOptional: false,
																	},
																},
																Node:       Example,
																IsOptional: false,
															},
														},
														Node:       ExampleItem,
														IsOptional: false,
													},
												},
												Node:       Examples,
												IsOptional: true,
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
														FullText:   "американский вариант английского языка; употребляется в США",
														Node:       Abbrev,
														Text:       "амер.",
														IsOptional: false,
													},
													&ArticleNode{
														IsItalics:  false,
														IsAccent:   false,
														Node:       Text,
														Text:       "; ",
														IsOptional: false,
													},
													&ArticleNode{
														FullText:   "разговорное",
														Node:       Abbrev,
														Text:       "разг.",
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
																Markup: []*ArticleNode{
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "доходное место",
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
																				Markup: []*ArticleNode{
																					&ArticleNode{
																						IsItalics:  false,
																						IsAccent:   false,
																						Node:       Text,
																						Text:       "plum job ",
																						IsOptional: false,
																					},
																					&ArticleNode{
																						IsItalics:  false,
																						IsAccent:   false,
																						Node:       Text,
																						Text:       "— тёплое местечко",
																						IsOptional: false,
																					},
																				},
																				Node:       Example,
																				IsOptional: false,
																			},
																		},
																		Node:       ExampleItem,
																		IsOptional: false,
																	},
																},
																Node:       Examples,
																IsOptional: true,
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
																		Text:       "выгодный заказ ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		Markup: []*ArticleNode{
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "(особенно предоставляемый за оказанную услугу)",
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
												},
												Node:       List,
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
														Text:       "тёмно-фиолетовый цвет ",
														IsOptional: false,
													},
													&ArticleNode{
														Markup: []*ArticleNode{
															&ArticleNode{
																IsItalics:  true,
																IsAccent:   false,
																Node:       Text,
																Text:       "(с оттенками бордового)",
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
																Markup: []*ArticleNode{
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "to wait for the plums to fall into one's mouth ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "— ждать, что сливы сами в рот посыплются; ждать, что поднесут всё на блюдечке",
																		IsOptional: false,
																	},
																},
																Node:       Example,
																IsOptional: false,
															},
														},
														Node:       ExampleItem,
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
																		Text:       "to speak with a plum in one's mouth ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		FullText:   "неодобрительное",
																		Node:       Abbrev,
																		Text:       "неодобр.",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  true,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "; ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		FullText:   "британский вариант английского языка; употребляется в Великобритании",
																		Node:       Abbrev,
																		Text:       "брит.",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "— иметь характерное для высших слоёв общества произношение",
																		IsOptional: false,
																	},
																},
																Node:       Example,
																IsOptional: false,
															},
														},
														Node:       ExampleItem,
														IsOptional: false,
													},
												},
												Node:       Examples,
												IsOptional: true,
											},
											&ArticleNode{
												Items: []*ArticleNode{
													&ArticleNode{
														Markup: []*ArticleNode{
															&ArticleNode{
																Dictionary: "LingvoUniversal (En-Ru)",
																ID:         "LingvoUniversal (En-Ru)__plum the plum-tree",
																Node:       CardRef,
																Text:       "plum the plum-tree",
																IsOptional: false,
															},
														},
														Node:       CardRefItem,
														IsOptional: false,
													},
												},
												Node:       CardRefs,
												IsOptional: false,
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
						Node:       ListItem,
						IsOptional: false,
					},
					&ArticleNode{
						Markup: []*ArticleNode{
							&ArticleNode{
								Markup: []*ArticleNode{
									&ArticleNode{
										FullText:   "имя прилагательное",
										Node:       Abbrev,
										Text:       "прил.",
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
												Markup: []*ArticleNode{
													&ArticleNode{
														FullText:   "устаревшее",
														Node:       Abbrev,
														Text:       "уст.",
														IsOptional: false,
													},
													&ArticleNode{
														IsItalics:  false,
														IsAccent:   false,
														Node:       Text,
														Text:       "полный, тучный",
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
																Markup: []*ArticleNode{
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "His face was fat and plum. ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "— У него было полное округлое лицо.",
																		IsOptional: false,
																	},
																},
																Node:       Example,
																IsOptional: false,
															},
														},
														Node:       ExampleItem,
														IsOptional: false,
													},
												},
												Node:       Examples,
												IsOptional: true,
											},
											&ArticleNode{
												Node:       Caption,
												Text:       "Syn:",
												IsOptional: true,
											},
											&ArticleNode{
												Markup: []*ArticleNode{
													&ArticleNode{
														Dictionary: "LingvoUniversal (En-Ru)",
														ID:         "LingvoUniversal (En-Ru)__fat",
														Node:       CardRef,
														Text:       "fat",
														IsOptional: false,
													},
												},
												Node:       Paragraph,
												IsOptional: true,
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
														FullText:   "диалектное",
														Node:       Abbrev,
														Text:       "диал.",
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
																Markup: []*ArticleNode{
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "мягкий ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		Markup: []*ArticleNode{
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "(о подушке)",
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
																Node:       Caption,
																Text:       "Syn:",
																IsOptional: true,
															},
															&ArticleNode{
																Markup: []*ArticleNode{
																	&ArticleNode{
																		Dictionary: "LingvoUniversal (En-Ru)",
																		ID:         "LingvoUniversal (En-Ru)__soft",
																		Node:       CardRef,
																		Text:       "soft",
																		IsOptional: false,
																	},
																},
																Node:       Paragraph,
																IsOptional: true,
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
																		Text:       "пышный, хорошо подходящий ",
																		IsOptional: false,
																	},
																	&ArticleNode{
																		Markup: []*ArticleNode{
																			&ArticleNode{
																				IsItalics:  true,
																				IsAccent:   false,
																				Node:       Text,
																				Text:       "(о тесте)",
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
																Node:       Caption,
																Text:       "Syn:",
																IsOptional: true,
															},
															&ArticleNode{
																Markup: []*ArticleNode{
																	&ArticleNode{
																		IsItalics:  false,
																		IsAccent:   false,
																		Node:       Text,
																		Text:       "well-raised",
																		IsOptional: false,
																	},
																},
																Node:       Paragraph,
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
														Text:       "ковкий, гибкий, податливый ",
														IsOptional: false,
													},
													&ArticleNode{
														Markup: []*ArticleNode{
															&ArticleNode{
																IsItalics:  true,
																IsAccent:   false,
																Node:       Text,
																Text:       "(о металле, камне)",
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
												Node:       Caption,
												Text:       "Syn:",
												IsOptional: true,
											},
											&ArticleNode{
												Markup: []*ArticleNode{
													&ArticleNode{
														Dictionary: "LingvoUniversal (En-Ru)",
														ID:         "LingvoUniversal (En-Ru)__soft",
														Node:       CardRef,
														Text:       "soft",
														IsOptional: false,
													},
												},
												Node:       Paragraph,
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
						Node:       ListItem,
						IsOptional: false,
					},
				},
				Node:       List,
				IsOptional: false,
			},
		},
	},
}

var testTranslationJSON = `
[
  {
    "Title": "plum",
    "TitleMarkup": [
      {
        "IsItalics": false,
        "IsAccent": false,
        "Node": "Text",
        "Text": "plum",
        "IsOptional": false
      }
    ],
    "Dictionary": "LingvoUniversal (En-Ru)",
    "ArticleId": "LingvoUniversal (En-Ru)__plum",
    "Body": [
      {
        "Markup": [
          {
            "Node": "Transcription",
            "Text": "plʌm",
            "IsOptional": false
          },
          {
            "IsItalics": false,
            "IsAccent": false,
            "Node": "Text",
            "Text": " ",
            "IsOptional": false
          },
          {
            "FullText": "британский вариант английского языка; употребляется в Великобритании",
            "Node": "Abbrev",
            "Text": "брит.",
            "IsOptional": false
          },
          {
            "IsItalics": false,
            "IsAccent": false,
            "Node": "Text",
            "Text": " ",
            "IsOptional": false
          },
          {
            "FileName": "plum.wav",
            "Node": "Sound",
            "Text": null,
            "IsOptional": false
          },
          {
            "IsItalics": false,
            "IsAccent": false,
            "Node": "Text",
            "Text": " / ",
            "IsOptional": false
          },
          {
            "FullText": "американский вариант английского языка; употребляется в США",
            "Node": "Abbrev",
            "Text": "амер.",
            "IsOptional": false
          },
          {
            "IsItalics": false,
            "IsAccent": false,
            "Node": "Text",
            "Text": " ",
            "IsOptional": false
          },
          {
            "FileName": "9612.wav",
            "Node": "Sound",
            "Text": null,
            "IsOptional": false
          }
        ],
        "Node": "Paragraph",
        "Text": null,
        "IsOptional": false
      },
      {
        "Type": 1,
        "Items": [
          {
            "Markup": [
              {
                "Markup": [
                  {
                    "FullText": "имя существительное",
                    "Node": "Abbrev",
                    "Text": "сущ.",
                    "IsOptional": false
                  }
                ],
                "Node": "Paragraph",
                "Text": null,
                "IsOptional": false
              },
              {
                "Type": 3,
                "Items": [
                  {
                    "Markup": [
                      {
                        "Type": 4,
                        "Items": [
                          {
                            "Markup": [
                              {
                                "Markup": [
                                  {
                                    "FullText": "ботаника",
                                    "Node": "Abbrev",
                                    "Text": "бот.",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "; = plum tree ",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "слива ",
                                    "IsOptional": false
                                  },
                                  {
                                    "Markup": [
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "(",
                                        "IsOptional": false
                                      },
                                      {
                                        "FullText": "латинский язык",
                                        "Node": "Abbrev",
                                        "Text": "лат.",
                                        "IsOptional": false
                                      },
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "Prunus; ",
                                        "IsOptional": false
                                      },
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "вид плодового дерева)",
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
                                    "Text": "слива ",
                                    "IsOptional": false
                                  },
                                  {
                                    "Markup": [
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "(плод)",
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
                                        "Dictionary": "LingvoUniversal (En-Ru)",
                                        "ArticleId": "LingvoUniversal (En-Ru)__French plum",
                                        "Node": "CardRef",
                                        "Text": "French plum",
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
                    ],
                    "Node": "ListItem",
                    "Text": null,
                    "IsOptional": false
                  },
                  {
                    "Markup": [
                      {
                        "Type": 4,
                        "Items": [
                          {
                            "Markup": [
                              {
                                "Markup": [
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "изюм",
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
                                    "Text": "сорт десертного винограда ",
                                    "IsOptional": false
                                  },
                                  {
                                    "Markup": [
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "(предназначенный для изготовления изюма)",
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
                                    "FullText": "американский вариант английского языка; употребляется в США",
                                    "Node": "Abbrev",
                                    "Text": "амер.",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "круглый леденец ",
                                    "IsOptional": false
                                  },
                                  {
                                    "Markup": [
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "(по форме напоминающий сливу)",
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
                                "Node": "Caption",
                                "Text": "Syn:",
                                "IsOptional": true
                              },
                              {
                                "Markup": [
                                  {
                                    "Dictionary": "LingvoUniversal (En-Ru)",
                                    "ArticleId": "LingvoUniversal (En-Ru)__sugarplum",
                                    "Node": "CardRef",
                                    "Text": "sugarplum",
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Paragraph",
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
                            "Text": "лакомый кусочек; нечто самое лучшее; сливки",
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
                                "Markup": [
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "to pick / take the plum ",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "— снимать сливки, отбирать самое лучшее",
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Example",
                                "Text": null,
                                "IsOptional": false
                              }
                            ],
                            "Node": "ExampleItem",
                            "Text": null,
                            "IsOptional": false
                          }
                        ],
                        "Node": "Examples",
                        "Text": null,
                        "IsOptional": true
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
                            "FullText": "американский вариант английского языка; употребляется в США",
                            "Node": "Abbrev",
                            "Text": "амер.",
                            "IsOptional": false
                          },
                          {
                            "IsItalics": false,
                            "IsAccent": false,
                            "Node": "Text",
                            "Text": "; ",
                            "IsOptional": false
                          },
                          {
                            "FullText": "разговорное",
                            "Node": "Abbrev",
                            "Text": "разг.",
                            "IsOptional": false
                          }
                        ],
                        "Node": "Paragraph",
                        "Text": null,
                        "IsOptional": false
                      },
                      {
                        "Type": 4,
                        "Items": [
                          {
                            "Markup": [
                              {
                                "Markup": [
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "доходное место",
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
                                        "Markup": [
                                          {
                                            "IsItalics": false,
                                            "IsAccent": false,
                                            "Node": "Text",
                                            "Text": "plum job ",
                                            "IsOptional": false
                                          },
                                          {
                                            "IsItalics": false,
                                            "IsAccent": false,
                                            "Node": "Text",
                                            "Text": "— тёплое местечко",
                                            "IsOptional": false
                                          }
                                        ],
                                        "Node": "Example",
                                        "Text": null,
                                        "IsOptional": false
                                      }
                                    ],
                                    "Node": "ExampleItem",
                                    "Text": null,
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Examples",
                                "Text": null,
                                "IsOptional": true
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
                                    "Text": "выгодный заказ ",
                                    "IsOptional": false
                                  },
                                  {
                                    "Markup": [
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "(особенно предоставляемый за оказанную услугу)",
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
                          }
                        ],
                        "Node": "List",
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
                            "Text": "тёмно-фиолетовый цвет ",
                            "IsOptional": false
                          },
                          {
                            "Markup": [
                              {
                                "IsItalics": true,
                                "IsAccent": false,
                                "Node": "Text",
                                "Text": "(с оттенками бордового)",
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
                                "Markup": [
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "to wait for the plums to fall into one's mouth ",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "— ждать, что сливы сами в рот посыплются; ждать, что поднесут всё на блюдечке",
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Example",
                                "Text": null,
                                "IsOptional": false
                              }
                            ],
                            "Node": "ExampleItem",
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
                                    "Text": "to speak with a plum in one's mouth ",
                                    "IsOptional": false
                                  },
                                  {
                                    "FullText": "неодобрительное",
                                    "Node": "Abbrev",
                                    "Text": "неодобр.",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": true,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "; ",
                                    "IsOptional": false
                                  },
                                  {
                                    "FullText": "британский вариант английского языка; употребляется в Великобритании",
                                    "Node": "Abbrev",
                                    "Text": "брит.",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "— иметь характерное для высших слоёв общества произношение",
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Example",
                                "Text": null,
                                "IsOptional": false
                              }
                            ],
                            "Node": "ExampleItem",
                            "Text": null,
                            "IsOptional": false
                          }
                        ],
                        "Node": "Examples",
                        "Text": null,
                        "IsOptional": true
                      },
                      {
                        "Type": null,
                        "Items": [
                          {
                            "Markup": [
                              {
                                "Dictionary": "LingvoUniversal (En-Ru)",
                                "ArticleId": "LingvoUniversal (En-Ru)__plum the plum-tree",
                                "Node": "CardRef",
                                "Text": "plum the plum-tree",
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
                        "IsOptional": false
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
                    "FullText": "имя прилагательное",
                    "Node": "Abbrev",
                    "Text": "прил.",
                    "IsOptional": false
                  }
                ],
                "Node": "Paragraph",
                "Text": null,
                "IsOptional": false
              },
              {
                "Type": 3,
                "Items": [
                  {
                    "Markup": [
                      {
                        "Markup": [
                          {
                            "FullText": "устаревшее",
                            "Node": "Abbrev",
                            "Text": "уст.",
                            "IsOptional": false
                          },
                          {
                            "IsItalics": false,
                            "IsAccent": false,
                            "Node": "Text",
                            "Text": "полный, тучный",
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
                                "Markup": [
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "His face was fat and plum. ",
                                    "IsOptional": false
                                  },
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "— У него было полное округлое лицо.",
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Example",
                                "Text": null,
                                "IsOptional": false
                              }
                            ],
                            "Node": "ExampleItem",
                            "Text": null,
                            "IsOptional": false
                          }
                        ],
                        "Node": "Examples",
                        "Text": null,
                        "IsOptional": true
                      },
                      {
                        "Node": "Caption",
                        "Text": "Syn:",
                        "IsOptional": true
                      },
                      {
                        "Markup": [
                          {
                            "Dictionary": "LingvoUniversal (En-Ru)",
                            "ArticleId": "LingvoUniversal (En-Ru)__fat",
                            "Node": "CardRef",
                            "Text": "fat",
                            "IsOptional": false
                          }
                        ],
                        "Node": "Paragraph",
                        "Text": null,
                        "IsOptional": true
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
                            "FullText": "диалектное",
                            "Node": "Abbrev",
                            "Text": "диал.",
                            "IsOptional": false
                          }
                        ],
                        "Node": "Paragraph",
                        "Text": null,
                        "IsOptional": false
                      },
                      {
                        "Type": 4,
                        "Items": [
                          {
                            "Markup": [
                              {
                                "Markup": [
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "мягкий ",
                                    "IsOptional": false
                                  },
                                  {
                                    "Markup": [
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "(о подушке)",
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
                                "Node": "Caption",
                                "Text": "Syn:",
                                "IsOptional": true
                              },
                              {
                                "Markup": [
                                  {
                                    "Dictionary": "LingvoUniversal (En-Ru)",
                                    "ArticleId": "LingvoUniversal (En-Ru)__soft",
                                    "Node": "CardRef",
                                    "Text": "soft",
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Paragraph",
                                "Text": null,
                                "IsOptional": true
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
                                    "Text": "пышный, хорошо подходящий ",
                                    "IsOptional": false
                                  },
                                  {
                                    "Markup": [
                                      {
                                        "IsItalics": true,
                                        "IsAccent": false,
                                        "Node": "Text",
                                        "Text": "(о тесте)",
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
                                "Node": "Caption",
                                "Text": "Syn:",
                                "IsOptional": true
                              },
                              {
                                "Markup": [
                                  {
                                    "IsItalics": false,
                                    "IsAccent": false,
                                    "Node": "Text",
                                    "Text": "well-raised",
                                    "IsOptional": false
                                  }
                                ],
                                "Node": "Paragraph",
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
                            "Text": "ковкий, гибкий, податливый ",
                            "IsOptional": false
                          },
                          {
                            "Markup": [
                              {
                                "IsItalics": true,
                                "IsAccent": false,
                                "Node": "Text",
                                "Text": "(о металле, камне)",
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
                        "Node": "Caption",
                        "Text": "Syn:",
                        "IsOptional": true
                      },
                      {
                        "Markup": [
                          {
                            "Dictionary": "LingvoUniversal (En-Ru)",
                            "ArticleId": "LingvoUniversal (En-Ru)__soft",
                            "Node": "CardRef",
                            "Text": "soft",
                            "IsOptional": false
                          }
                        ],
                        "Node": "Paragraph",
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
]
`

BEGIN TRANSACTION;

DROP TABLE IF EXISTS "word" CASCADE;
DROP SEQUENCE IF EXISTS "word_seq";

---------------------------------------------------------

CREATE SEQUENCE "word_seq";
CREATE TABLE "word" (

  "id"           BIGINT PRIMARY KEY       DEFAULT "nextval"('"word_seq"'),
  "created_date" TIMESTAMP WITH TIME ZONE DEFAULT "now"(),
  "value"        TEXT NOT NULL,
  "lang_from"    TEXT NOT NULL,
  "lang_to"      TEXT NOT NULL,
  "translate"    JSONВ NOT NULL
);

END TRANSACTION;

INSERT INTO word (value, lang_from, lang_to, translate) VALUES ('house', 'en', 'ru', '{
  "head": {},
  "def": [
    {
      "text": "house",
      "pos": "noun",
      "ts": "haʊs",
      "tr": [
        {
          "text": "дом",
          "pos": "noun",
          "gen": "м",
          "syn": [
            {
              "text": "здание",
              "pos": "noun",
              "gen": "ср"
            },
            {
              "text": "домик",
              "pos": "noun",
              "gen": "м"
            },
            {
              "text": "жилой дом",
              "pos": "noun"
            },
            {
              "text": "домишко",
              "pos": "noun",
              "gen": "м"
            }
          ],
          "mean": [
            {
              "text": "building"
            },
            {
              "text": "cottage"
            },
            {
              "text": "apartment house"
            }
          ],
          "ex": [
            {
              "text": "auction house",
              "tr": [
                {
                  "text": "аукционный дом"
                }
              ]
            },
            {
              "text": "father''s house",
              "tr": [
                {
                  "text": "отчий дом"
                }
              ]
            },
            {
              "text": "single storey house",
              "tr": [
                {
                  "text": "одноэтажный дом"
                }
              ]
            },
            {
              "text": "old wooden house",
              "tr": [
                {
                  "text": "старый деревянный дом"
                }
              ]
            },
            {
              "text": "build new houses",
              "tr": [
                {
                  "text": "строить новые дома"
                }
              ]
            },
            {
              "text": "timber frame house",
              "tr": [
                {
                  "text": "каркасный дом"
                }
              ]
            },
            {
              "text": "own fashion house",
              "tr": [
                {
                  "text": "собственный модный дом"
                }
              ]
            },
            {
              "text": "dwelling house",
              "tr": [
                {
                  "text": "жилое здание"
                }
              ]
            },
            {
              "text": "summer house",
              "tr": [
                {
                  "text": "летний домик"
                }
              ]
            }
          ]
        },
        {
          "text": "палата представителей",
          "pos": "noun",
          "syn": [
            {
              "text": "палата",
              "pos": "noun",
              "gen": "ж"
            },
            {
              "text": "палата парламента",
              "pos": "noun"
            }
          ],
          "mean": [
            {
              "text": "house of representatives"
            },
            {
              "text": "chamber"
            },
            {
              "text": "house of parliament"
            }
          ],
          "ex": [
            {
              "text": "lower house of parliament",
              "tr": [
                {
                  "text": "нижняя палата российского парламента"
                }
              ]
            }
          ]
        },
        {
          "text": "жилье",
          "pos": "noun",
          "gen": "ср",
          "syn": [
            {
              "text": "жилище",
              "pos": "noun",
              "gen": "ср"
            }
          ],
          "mean": [
            {
              "text": "housing"
            },
            {
              "text": "home"
            }
          ],
          "ex": [
            {
              "text": "safe houses",
              "tr": [
                {
                  "text": "безопасное жилье"
                }
              ]
            },
            {
              "text": "right to housing",
              "tr": [
                {
                  "text": "право на жилище"
                }
              ]
            }
          ]
        },
        {
          "text": "хаус",
          "pos": "noun",
          "gen": "м",
          "mean": [
            {
              "text": "haus"
            }
          ],
          "ex": [
            {
              "text": "coffee house",
              "tr": [
                {
                  "text": "кофе хаус"
                }
              ]
            }
          ]
        },
        {
          "text": "квартира",
          "pos": "noun",
          "gen": "ж",
          "mean": [
            {
              "text": "apartment"
            }
          ]
        },
        {
          "text": "изба",
          "pos": "noun",
          "gen": "ж",
          "syn": [
            {
              "text": "хата",
              "pos": "noun",
              "gen": "ж"
            }
          ],
          "mean": [
            {
              "text": "hut"
            }
          ]
        }
      ]
    },
    {
      "text": "house",
      "pos": "adjective",
      "ts": "haʊs",
      "tr": [
        {
          "text": "домашний",
          "pos": "adjective",
          "syn": [
            {
              "text": "домовой",
              "pos": "adjective"
            }
          ],
          "mean": [
            {
              "text": "home"
            }
          ],
          "ex": [
            {
              "text": "house cat",
              "tr": [
                {
                  "text": "домашняя кошка"
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "text": "house",
      "pos": "verb",
      "ts": "haʊs",
      "tr": [
        {
          "text": "размещаться",
          "pos": "verb",
          "asp": "несов",
          "syn": [
            {
              "text": "разместиться",
              "pos": "verb",
              "asp": "сов"
            }
          ],
          "mean": [
            {
              "text": "host"
            },
            {
              "text": "locate"
            }
          ]
        },
        {
          "text": "расквартировывать",
          "pos": "verb",
          "asp": "несов"
        }
      ]
    }
  ]
}');
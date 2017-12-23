package yandex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const baseURLPath = "/translate"

func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))

	server := httptest.NewServer(apiHandler)

	// client being tested and is configured to use test server.
	client = NewClient(server.URL + baseURLPath + "/")

	return client, mux, server.Close
}

func TestTranslateSuccess(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	dictionary := &Dictionary{client, "token"}

	mux.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
						"head": {},
						"def": [
							{
								"text": "calc",
								"pos": "noun",
								"ts": "kælk",
								"tr": [
									{
										"text": "программа calc",
										"pos": "noun"
									},
									{
										"text": "калькулятор",
										"pos": "noun",
										"gen": "м",
										"mean": [
											{
												"text": "calculator"
											}
										]
									}
								]
							}
						]
					}`)
	})

	word, err := dictionary.Translate("calc", "en", "ru")
	if err != nil {
		t.Errorf("Dictionary.Translate returned error: %v", err)
	}

	if word.Def[0].Text != "calc" {
		t.Errorf("Dictionary.translate returned %+v", word)
	}
}

func TestTranslateSuccessEmptyResult(t *testing.T) {
	client, mux, teardown := setup()

	defer teardown()

	mux.HandleFunc("/lookup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{
						"head": {},
						"def": []
					}`)
	})

	dictionary := &Dictionary{client, "token"}

	word, err := dictionary.Translate("calc", "en", "ru")
	if err != nil {
		t.Errorf("Dictionary.Translate returned error: %v", err)
	}

	if len(word.Def) != 0 {
		t.Errorf("Dictionary.translate returned %+v", word)
	}
}

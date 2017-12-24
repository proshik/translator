package handler

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/proshik/translator/yandex"
	"encoding/json"
)

type DictionaryMock struct {
	response *yandex.WordResponse
	error *yandex.ExternalError
}

func (dm *DictionaryMock) Translate(text, langFrom, langTo string) (*yandex.WordResponse, *yandex.ExternalError) {
	return dm.response, dm.error
}

func TestWord(t *testing.T) {
	var wr yandex.WordResponse
	json.Unmarshal([]byte(`{
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
					}`), &wr)

	h := &Handler{&DictionaryMock{&wr, nil}}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/lookup?text=calc&lang=en-ru", nil)

	h.Word(w, r)

	resp := w.Result()
	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}
}

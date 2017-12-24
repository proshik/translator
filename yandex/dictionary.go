package yandex

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

const defaultBaseURL = "https://dictionary.yandex.net/api/v1/dicservice.json/"

//Translate response from Yandex dictionary
type WordResponse struct {
	Head Head  `json:"head"`
	Def  []Def `json:"def"`
}

type Head struct{}

type Def struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
	Ts   string `json:"ts"`
	Tr   []Tr   `json:"tr"`
}

type Ex struct {
	Attr
	Tr
}

type Tr struct {
	Attr
	Syn  []Syn  `json:"syn"`
	Mean []Mean `json:"mean"`
	Ex   []Ex   `json:"ex"`
}

type Syn struct {
	Attr
}

type Mean struct {
	Attr
}

type Attr struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
	Gen  string `json:"gen"`
}

//Error from Yandex dictionary
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type DictionaryTranslator interface {
	Translate(text, langFrom, langTo string) (*WordResponse, *ExternalError)
}

type Dictionary struct {
	client *Client
	token  string
}

func NewDictionary(token string) DictionaryTranslator {
	return &Dictionary{
		NewClient(defaultBaseURL),
		token,
	}
}

func (dict *Dictionary) Translate(text, langFrom, langTo string) (*WordResponse, *ExternalError) {
	url := fmt.Sprintf("lookup?lang=%s-%s&key=%s&text=%s", langFrom, langTo, dict.token, text)

	u, err := dict.client.BaseURL.Parse(url)
	if err != nil {
		return nil, &ExternalError{err, UNEXPECTED_ERROR}
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, &ExternalError{err, UNEXPECTED_ERROR}
	}

	resp, err := dict.client.Do(req)
	if err != nil {
		return nil, &ExternalError{err, UNEXPECTED_ERROR}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		d := json.NewDecoder(resp.Body)

		var erResp ErrorResponse
		err = d.Decode(&erResp)
		if err != nil {
			return nil, &ExternalError{err, UNEXPECTED_ERROR}
		}

		return nil, &ExternalError{errors.New("Error from external system"), toCode(erResp.Code)}
	}

	d := json.NewDecoder(resp.Body)

	var tr WordResponse
	err = d.Decode(&tr)
	if err != nil {
		return nil, &ExternalError{err, UNEXPECTED_ERROR}
	}

	return &tr, nil
}

func toCode(code int) Code {
	switch code {
	case 401:
		return KEY_INVALID
	case 402:
		return KEY_BLOCKED
	case 403:
		return DAILY_REQ_LIMIT_EXCEEDED
	case 413:
		return TEXT_TOO_LONG
	case 501:
		return LANG_NOT_SUPPORTED
	default:
		return UNEXPECTED_ERROR
	}
}

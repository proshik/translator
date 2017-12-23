package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/proshik/translator/yandex"
	"log"
	"net/http"
	"strings"
)

//LangDirection describe a direction of translate(en-ru, ru-en, etc)
type LangDirection struct {
	LangFrom string
	LangTo   string
}

type Handler struct {
	dictionary *yandex.Dictionary
}

func NewHandler(dict *yandex.Dictionary) *Handler {
	return &Handler{dict}
}

func (handler *Handler) Word(w http.ResponseWriter, r *http.Request) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		http.Error(w, "parameter lang is empty", http.StatusBadRequest)
		return
	}

	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "parameter text is empty", http.StatusBadRequest)
		return
	}

	langDirection, err := extractLangDirection(lang)
	if err != nil {
		log.Print(err)
		http.Error(w, "Incorrect lang value", http.StatusBadRequest)
		return
	}

	translate, extErr := handler.dictionary.Translate(text, langDirection.LangFrom, langDirection.LangTo)
	if extErr != nil && extErr.Err != nil {
		fmt.Println(extErr)
		http.Error(w, "Error to translate a text", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(translate.Def[0]); err != nil {
		log.Print(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) Text(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func extractLangDirection(lang string) (*LangDirection, error) {
	result := strings.Split(lang, "-")
	if len(result) != 2 {
		return nil, errors.New(fmt.Sprintf("Error on spit lang value=%s", lang))
	}

	return &LangDirection{result[0], result[1]}, nil
}

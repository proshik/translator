package main

import (
	"net/http"
	"log"
	"github.com/go-chi/chi"
	"github.com/proshik/translator/config"
	"github.com/proshik/translator/yandex"
	"github.com/proshik/translator/handler"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	dict := yandex.NewDictionary(cfg.YandexDictionaryToken)

	h := handler.NewHandler(dict)

	r := chi.NewRouter()
	r.Route("/translate", func(r chi.Router) {
		r.Get("/word", h.Word)

		r.Post("/text", h.Text)
	})

	log.Println("Api is waiting for requests...")

	http.ListenAndServe(":8080", r)
}

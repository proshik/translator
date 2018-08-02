package main

import (
	"github.com/go-chi/chi"
	"github.com/proshik/translator/config"
	"github.com/proshik/translator/handler"
	"github.com/proshik/translator/yandex"
	"log"
	"net/http"
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
	})

	log.Println("Api is waiting for requests...")

	panic(http.ListenAndServe(":"+cfg.Port, r))
}

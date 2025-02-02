package faas

import (
	"net/http"

	"github.com/jeffleon2/shipping-go-hello-api/handlers/rest"
	"github.com/jeffleon2/shipping-go-hello-api/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translationHandler := rest.NewTranslateHandler(translationService)
	translationHandler.TranslateHandler(w, r)
}

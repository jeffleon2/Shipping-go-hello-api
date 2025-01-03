// Package faas wrapper for cloud functions.
package faas

import (
	"net/http"

	"github.com/jeffleon2/shipping-go-hello-api/handlers/rest"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}

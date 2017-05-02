package web

import (
	"net/http"
)

// Redirect is for http redirects but its using types from web package
func Redirect(response http.ResponseWriter, request *http.Request, location Path) {
	http.Redirect(response, request, location.String(), http.StatusFound)
}

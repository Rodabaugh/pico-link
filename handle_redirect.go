package main

import (
	"net/http"
)

func handlerRedirect(w http.ResponseWriter, r *http.Request) {
	newUrl := "https://google.com/"
	http.Redirect(w, r, newUrl, http.StatusSeeOther)
}

package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerRedirect(w http.ResponseWriter, r *http.Request) {
	linkName := r.PathValue("link_name")

	linkUrl, err := cfg.db.GetLinkByName(r.Context(), linkName)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Unable to find link", err)
	}

	http.Redirect(w, r, linkUrl.LinkUrl, http.StatusSeeOther)
}

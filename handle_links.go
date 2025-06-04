package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Rodabaugh/pico-link/internal/database"
	"github.com/google/uuid"
)

type Link struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	LinkName  string
	LinkUrl   string
}

func (cfg *apiConfig) handlerCreateLink(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		LinkName string `json:"link_name"`
		LinkURL  string `json:"link_url"`
	}

	type response struct {
		Link
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Was unable to decode parameters", err)
		return
	}

	link, err := cfg.db.CreateLink(r.Context(), database.CreateLinkParams{
		LinkName: params.LinkName,
		LinkUrl:  params.LinkURL,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to create Link", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		Link: Link{
			ID:        link.ID,
			CreatedAt: link.CreatedAt,
			UpdatedAt: link.UpdatedAt,
			LinkName:  link.LinkName,
			LinkUrl:   link.LinkUrl,
		},
	})
}

func (cfg *apiConfig) handlerGetAllLinks(w http.ResponseWriter, r *http.Request) {
	links, err := cfg.Links()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to get links", err)
		return
	}

	respondWithJSON(w, http.StatusOK, links)
}

func (cfg *apiConfig) Links() ([]Link, error) {
	databaseLinks, err := cfg.db.GetAllLinks(context.Background())
	if err != nil {
		return nil, fmt.Errorf("unable to get links from database", err)
	}

	links := []Link{}

	for _, dbLink := range databaseLinks {
		links = append(links, Link{
			ID:        dbLink.ID,
			CreatedAt: dbLink.CreatedAt,
			UpdatedAt: dbLink.UpdatedAt,
			LinkName:  dbLink.LinkName,
			LinkUrl:   dbLink.LinkUrl,
		})
	}

	return links, nil
}

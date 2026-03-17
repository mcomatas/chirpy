package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/mcomatas/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUpgradeToChirpyRed(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	api_key, err := auth.GetAPIKey(r.Header)
	if err != nil || api_key != cfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "invalid API key", err)
		return
	}

	params := parameters{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = cfg.db.UpgradeToChirpyRed(r.Context(), params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "user not found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

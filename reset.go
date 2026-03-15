package main

import "net/http"

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Forbidden", nil)
		return
	}

	cfg.fileserverHits.Store(0)

	if err := cfg.db.DeleteAllUsers(r.Context()); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not delete users", err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reset OK"))
}

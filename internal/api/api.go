package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sallescosta/ama/internal/store/pgstore"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	return apiHandler{q: q, r: chi.NewRouter()}
}

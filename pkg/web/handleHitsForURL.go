package web

import(
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"database/sql"
)

type HitsResponse struct{
	Hits int `json:"hits"`
	RequestKey string `json:"key"`
}
func handleHitsForURL(ctx *Context, rw http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	targetKey := vars["key"]
	hits, err := ctx.Redis.GetHits(targetKey)
	if err != nil {
		if err == sql.ErrNoRows{
			return http.StatusBadRequest, err
		}
		return http.StatusInternalServerError, err
	}
	hr := HitsResponse{Hits: hits, RequestKey: targetKey}
	bytes, err := json.Marshal(hr)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = rw.Write(bytes)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

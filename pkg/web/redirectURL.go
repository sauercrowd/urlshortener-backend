package web

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRedirectURL(ctx *Context, rw http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	targetKey := vars["key"]
	url, err := ctx.Redis.GetURL(targetKey)
	if err != nil && err != sql.ErrNoRows { //real error
		log.Println("Could not get key from redis:", err)
		return http.StatusInternalServerError, err
	}
	if err == sql.ErrNoRows {
		log.Println("Could not find key")
		return showKeyNotFound(targetKey, rw)
	}
	rw.Header().Add("Cache-Control","no-store, no-cache, must-revalidate")
	http.Redirect(rw, r, url, http.StatusFound)
	log.Println("Redirect to:", url)
	return http.StatusFound, nil
}

func showKeyNotFound(key string, rw http.ResponseWriter) (int, error) {
	return http.StatusNotFound, nil
}

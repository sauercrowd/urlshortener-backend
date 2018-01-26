package web

import (
	"encoding/json"
	"net/http"
)

func handleListURLs(ctx *Context, rw http.ResponseWriter, r *http.Request) (int, error) {
	urls, err := ctx.Redis.ListURLs()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	bytes, err := json.Marshal(urls)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = rw.Write(bytes)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

package web

import (
	"net/url"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type createURLRequest struct {
	URL string `json:"url"`
}

type createURLResponse struct {
	Short string `json:"key"`
}

func handleCreateURL(ctx *Context, rw http.ResponseWriter, r *http.Request) (int, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Println("Could not read body", err)
		return http.StatusInternalServerError, err
	}
	var urlRequest createURLRequest
	if err := json.Unmarshal(bytes, &urlRequest); err != nil {
		log.Println("Could not unmarshal body into struct:", err)
		return http.StatusBadRequest, err
	}

	//check if it is a valid URL
	_, err = url.ParseRequestURI(urlRequest.URL)
	if err != nil{
		log.Println("Not a valid URL: ",urlRequest.URL)
		return http.StatusBadRequest, err
	}

	short, err := ctx.Redis.AddURL(urlRequest.URL)
	if err != nil {
		log.Println("Could not add URL to redis", err)
		return http.StatusInternalServerError, err
	}
	urlResponse := createURLResponse{Short: createShort(ctx.Basename, short)}
	return sendJSON(rw, urlResponse)
}

func createShort(basename string, key string) string {
	return key
}

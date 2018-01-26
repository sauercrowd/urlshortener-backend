package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sauercrowd/urlshortener-backend/pkg/persistence"
)

// Context gives additional context to webhandlers
type Context struct {
	Redis    *persistence.Storage
	Basename string
}

// Handler handles webhandlers with additional context
type Handler struct {
	*Context
	Handle func(*Context, http.ResponseWriter, *http.Request) (int, error)
}

func (wh Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := wh.Handle(wh.Context, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}

func Handle(addr string, ctx *Context) {
	r := mux.NewRouter()

	r.Handle("/{key:[A-Za-z]+}", Handler{ctx, handleRedirectURL}).Methods("GET")
	r.Handle("/{key:[A-Za-z]+}/hits", Handler{ctx, handleHitsForURL}).Methods("GET")
	r.Handle("/api/v1/add", Handler{ctx, handleCreateURL}).Methods("POST")
	r.Handle("/api/v1/list", Handler{ctx, handleListURLs}).Methods("GET")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("frontend"))))

	log.Fatal(http.ListenAndServe(addr, r))
}

func sendJSON(rw http.ResponseWriter, data interface{}) (int, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Println("Could not serialize response:", err)
		return http.StatusInternalServerError, err
	}
	rw.Header().Add("Content-Type", "application/json")
	_, err = rw.Write(bytes)
	if err != nil {
		log.Println("Could not write response body:", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

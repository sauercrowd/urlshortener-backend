package persistence

import (
	"database/sql"
	"log"

)

func (r *Storage) GetURL(key string) (string, error) {
	var url string
	err := r.client.QueryRow(`SELECT url FROM urls WHERE key=$1`,key).Scan(&url)
	if err == nil {
		addHit(r.client, key)
	}
	return url, err
}

func addHit(client *sql.DB, key string) {
	if err := client.QueryRow(`UPDATE urls SET hits = hits + 1 WHERE key = $1`,key).Scan(); err != nil && err != sql.ErrNoRows{
		log.Println("Error: ", err)
	}
}

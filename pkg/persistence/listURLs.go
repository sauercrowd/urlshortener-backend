package persistence

import (
	"log"
)

type ShortURL struct {
	Hits      int    `json:"hits"`
	ShortKey  string `json:"key"`
	TargetURL string `json:"target"`
}

func (r *Storage) ListURLs() ([]ShortURL, error) {
	ret := make([]ShortURL, 0)
	//get keys
	rows, err := r.client.Query(`SELECT url, key, hits FROM urls ORDER BY hits DESC LIMIT 10`)
	if err != nil {
		return nil, err
	}
	for rows.Next(){
		var su ShortURL
		if err := rows.Scan(&su.TargetURL, &su.ShortKey, &su.Hits); err != nil{
			log.Println("Error, could not get informations from line", err)
			return nil, err
		}
		ret = append(ret, su)
	}
	return ret, nil
}

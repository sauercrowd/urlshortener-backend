package persistence

import (
	"database/sql"
	"fmt"
	"log"
)

func (r *Storage) AddURL(url string) (string, error) {
	return r.generateKey(url, 0)
}

const maxtries = 3
const criticaltries = 9

func (r *Storage) generateKey(url string, try int) (string, error) {
	if try > criticaltries{
		log.Println("Exceeded tries")
		return "",fmt.Errorf("Exceeded tries")
	}
	if try % maxtries == maxtries-1 { // every maxtries increase the count
		err := r.SetCharacterCount(r.characterCount+1)
		if err != nil{
			return "", err
		}
		//continue 
	}
	//generate a random id
	key, err := generateID(r.characterCount)
	if err != nil {
		return "", err
	}
	//check if the key exists
	var count int
	if err := r.client.QueryRow(`SELECT COUNT(key) FROM urls WHERE key=$1`,key).Scan(&count); err != nil{
		log.Println("Could not get count for key")
		return "", err
	}
	//key found, try again
	if count > 0 {
		return r.generateKey(url, try+1)
	}
	//key is ok, set it
	if err := r.client.QueryRow(`INSERT INTO urls(url, key, hits) VALUES($1, $2, 0)`,url, key).Scan(); err != nil && err != sql.ErrNoRows{
		log.Println("Could not insert into table", err)
		return "", err
	}
	return key, nil
}

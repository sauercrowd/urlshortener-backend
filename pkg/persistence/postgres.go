package persistence

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
	"github.com/sauercrowd/urlshortener-backend/pkg/flags"
)

type Storage struct {
	client *sql.DB
	characterCount int
}
const DBNAME = "short"

func Create(f *flags.Flags) (*Storage, error) {
	dbstr :=fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", f.Puser, f.Ppassword, f.Phost, 5432, f.Puser)
	conn, err := sql.Open("postgres", dbstr)
	if err != nil {
		return nil, fmt.Errorf("Could not ping redis: %v", err)
	}
	if err := createDBIfNotExists(conn, DBNAME); err != nil{
		return nil, err
	}
	if err := conn.Close(); err != nil {
		return nil, err
	}
	dbstr = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", f.Puser, f.Ppassword, f.Phost, 5432, DBNAME)
	conn, err = sql.Open("postgres", dbstr)
	if err := createSidebarTableIfNotExist(conn); err != nil {
		return nil, err
	}
	s := Storage{client: conn, characterCount: 1}
	//TODO!
	// err = r.getOrCreateCharacterCount()
	// if err != nil{
	// 	return nil, err
	// }
	return &s, nil
}

func createDBIfNotExists(db *sql.DB, name string) error {
	var count int64
	err := db.QueryRow("SELECT COUNT(1) FROM pg_database WHERE datname = $1", name).Scan(&count)
	//return if database exists or error happend
	if err != nil || count == 1 {
		if err == nil {
			err = db.Close()
		}
		return err
	}
	err = db.QueryRow(fmt.Sprintf("CREATE DATABASE %s", name)).Scan()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	return nil
}

func createSidebarTableIfNotExist(db *sql.DB) error {
	err := db.QueryRow("CREATE TABLE IF NOT EXISTS urls(url text, key text PRIMARY KEY, hits int)").Scan()
	if err == sql.ErrNoRows {
		return nil
	}
	return err
}



//COUNTKEY is the key used by redis to save how many characters are currently used for short generation
const COUNTKEY="CHARCOUNT"

// func (r *Storage) getOrCreateCharacterCount() error {
// 	keyCount, err := r.client.Exists(COUNTKEY).Result()
// 	if err != nil{
// 		log.Println("Could not get Character count")
// 		return err
// 	}
// 	//if there's nothing saved, return 1 to start with 1 character
// 	if keyCount == 0{
// 		return r.SetCharacterCount(1)
// 	}
// 	//since nothing is found, get the count from redis
// 	countStr, err := r.client.Get(COUNTKEY).Result()
// 	if err != nil{
// 		log.Println("Could not get character count:")
// 		log.Println(err)
// 		return err
// 	}
// 	count, err := strconv.Atoi(countStr)
// 	if err != nil{
// 		log.Println("Could not convert character count string to int")
// 		log.Println(err)
// 		return err
// 	}
// 	r.characterCount = count
// 	return nil
// }

func (r *Storage) SetCharacterCount(count int) error {
	//TODO
	// _, err := r.client.Set(COUNTKEY, count,0).Result()
	// if err != nil{
	// 	return err
	// }
	r.characterCount=count
	return nil
}
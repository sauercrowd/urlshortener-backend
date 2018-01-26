package persistence

func (r *Storage) GetHits(key string) (int, error) {
	var hits int
	err := r.client.QueryRow(`SELECT hits FROM urls WHERE key=$1`,key).Scan(&hits)
	if err != nil{
		return -1, err
	}
	return hits, err
}
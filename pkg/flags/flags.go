package flags

import (
	"flag"
)

// command line flags
type Flags struct {
	RedisHost, RedisPassword string
	RedisDB                  int
	Port                     int
}

func Parse() *Flags {
	var f Flags
	flag.StringVar(&f.RedisHost, "rhost", "127.0.0.1", "Redis Host")
	flag.StringVar(&f.RedisPassword, "rpass", "", "Redis Password")
	flag.IntVar(&f.RedisDB, "rdb", 0, "redis db")
	flag.IntVar(&f.Port, "p", 8080, "port to listen on")
	flag.Parse()
	return &f
}

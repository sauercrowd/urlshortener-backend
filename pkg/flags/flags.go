package flags

import (
	"flag"
)

// command line flags
type Flags struct {
	Phost, Ppassword, Puser string
	Port                     int
}

func Parse() *Flags {
	var f Flags
	flag.StringVar(&f.Phost, "phost", "127.0.0.1", "Postgres Host")
	flag.StringVar(&f.Ppassword, "ppass", "postgres", "Postgres Password")
	flag.StringVar(&f.Puser, "puser", "postgres", "Postgres User")
	flag.IntVar(&f.Port, "p", 8080, "port to listen on")
	flag.Parse()
	return &f
}

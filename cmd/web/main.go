package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// log...
type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	// log...
	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorlog: errorlog,
		infolog:  infolog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.routes(),
	}

	// log.Printf("Statring server on %s", *addr)

	infolog.Printf("Statring server on %s", *addr)
	// err := http.ListenAndServe(*addr, mux)

	err := srv.ListenAndServe()
	errorlog.Fatal(err)

}

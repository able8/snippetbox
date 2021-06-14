package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/able8/snippetbox/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

// Add a snippets field to the application struct.
// This will allow us to make the SnippetModel objecct available to our handles.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "MySQL data source name")
	secret := flag.String("secret", "asdfaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "Secret key, 32 bytes long")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// Initialize a tls.Config struct to hold the non-default TLS settings.
	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS13,
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	srv := &http.Server{
		Addr:      *addr,
		ErrorLog:  errorLog,
		Handler:   app.routes(),
		TLSConfig: tlsConfig,
		// Add Idle, Read adn Write timeouts to the server.
		// We set IdleTimeout to 1 minute, which means that all keep-alive connections will
		// be automatically closed after 1 minute of inactivity.
		IdleTimeout: time.Minute,
		// If the request headers or body are still being read  5s after the request is first accepted,
		// then Go will close the underlying connection. Because this is a hard closure on the connection,
		// the user won't receive any HTTP(s) response.
		// Setting a short ReadTimeout period helps to mitigate the risk from slow-client attacks.
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		// Control the maximum number of bytes the server will read when parsing request headers.
		// By default, Go allows a maximum header lenght of 1MB.
		// MaxHeaderBytes: 900,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// err = srv.ListenAndServe()
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	log.Printf("DB connection established")
	return db, nil
}

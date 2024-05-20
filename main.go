package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/anthoturc/simple-token-service/templates"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const CSRF_STRING = "asdfj4qqaBASBQxsa"

func main() {

	db, err := OpenSqlDb()
	if err != nil {
		log.Fatalf("failed to open db connection with connstr: %s", err.Error())
	}
	defer db.Close()

	dbService := &DbService{
		DB: db,
	}

	authMw := AuthMiddleware{
		db: dbService,
	}

	r := chi.NewRouter()

	cp := ControlPlane{
		db: dbService,
	}
	cp.Templates.GetToken = Must(ParseFS(templates.FS, "get_token.gohtml", "base.gohtml"))
	cp.Templates.Token = Must(ParseFS(templates.FS, "token.gohtml", "base.gohtml"))

	csrfMw := csrf.Protect(
		[]byte(CSRF_STRING),
		csrf.Secure(false),
	)

	r.Route("/", func(r chi.Router) {
		r.Use(csrfMw)
		r.Get("/", cp.Home)
		r.Post("/token", cp.Token)
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(authMw.Authenticate)
		r.Get("/hello", Hello)
	})

	log.Fatal(http.ListenAndServe(":8000", r))
}

func OpenSqlDb() (*sql.DB, error) {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=dev password=dev dbname=sts sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	return db, nil
}

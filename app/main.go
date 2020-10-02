package main

import (
    "fmt"
    // "log"
    "net/http"
    // "github.com/gorilla/mux"
    // "github.com/golang-migrate/migrate/v4"
    // _ "github.com/lib/pq"
    // "database/sql"

)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is the home page. Welcome!")
}

func main() {
    // _, err := migrate.New(
    //     "github://mattes:personal-access-token@mattes/migrate_test",
    //     "postgres://localhost:5432/database?sslmode=enable")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // _, err = postgres.WithInstance(db, &postgres.Config{})
    // m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)
    // m.Steps(2)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // r := mux.NewRouter()
    // r.HandleFunc("/", homePage)
}


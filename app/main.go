package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

type constError string

var (
	UserNotFound = constError("UserNotFound")
)

type User struct {
	Id        string
	Name      string
	CreatedAt time.Time
}

func NewUser(name string) (user User, err error) {
	userId, err := uuid.NewRandom()
	if err != nil {
		return user, err
	}
	user = User{Id: userId.String(), Name: name, CreatedAt: time.Now()}
	return user, err
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page. Welcome!")
}

func insertUser(db *sql.DB, user *User) error {
	stmt, err := db.Prepare("INSERT INTO users (id, name, created_at) VALUES($1, $2, $3)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Id, user.Name, user.CreatedAt)
	if err != nil {
		return err
	}
	return err
}

func rowToUser(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.Id, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func fetchUser(db *sql.DB, userId string) (*User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", userId)
	return rowToUser(row)
}

func main() {
	dbstring, ok := os.LookupEnv("DBSTRING!")
	if !ok {
		dbstring = "host=db user=app password=app sslmode=disable"
	}
	db, err := sql.Open("postgres", dbstring)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	newUser, err := NewUser("test")
	if err != nil {
		log.Fatal(err)
	}
	err = insertUser(db, &newUser)
	if err != nil {
		fmt.Println("insert fail")
		log.Fatal(err)
	}
	fmt.Println("inserted")
	queriedUser, err := fetchUser(db, newUser.Id)
	if err != nil {
		fmt.Println("query fail")
		log.Fatal(err)
	}
	fmt.Println("queried")
	fmt.Println(queriedUser)
}

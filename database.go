package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func Database(usage string) *sql.DB {
	/*
		THIS OPENS THE DATABASE CONNECTION. NOTE THAT
		THE DATABASE IS BASICALLY IN WAIT MODE, THE
		CONNECTION ONLY ACTUALLY OPENS WHEN A QUERY IS
		MADE.
	*/
	dbInfo := "dbname=socialmediasite sslmode=disable"
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	// THIS CONFIRMS DATABASE IS OPEN-ABLE
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Database connection established. "+
		"Attempting to %s.", usage))
	return db
}

func Encrypt(password string) string {
	h := sha256.New()
	_, err := io.WriteString(h, password)
	if err != nil {
		log.Fatal(err, "Unknown hashing error.")
	}
	return hex.EncodeToString([]byte(fmt.Sprint(h)))
}

func AddNewUserAccount(age int, firstname string, lastname string,
	email string, gender string, public bool, password string) {
	/*
		THIS CONNECTS TO THE DATABASE AND ADDS A USER
	*/
	db := Database("add user")
	query := fmt.Sprintf("INSERT INTO user_account(age, firstname, lastname, email, "+
		"gender, public, joindate, active, password)"+
		"VALUES (%d, '%s', '%s', '%s', '%s', '%t', now(), true, '%s');",
		age, firstname, lastname, email, gender, public, Encrypt(password))
	_, err := db.Query(query)
	if err != nil {
		DatabaseErrorHandler(err)
	}
	log.Println(fmt.Sprintf("Successfully added user <%s> to Database.", email))

	// EVERY DATABASE USAGE MUST FINISH WITH THE DATABASE BEING CLOSED
	defer db.Close()
}

type user struct {
	id        int
	firstname string
	lastname  string
	email     string
	gender    string
	public    bool
	joindate  string
	active    bool
}

func PrintUser(u user) {
	/*
	THIS IS A DEBUGGING TOOL
	 */
	log.Printf("\n\n\tUSER {\n" +
			"\tId: %v\n" +
			"\tFirst Name: %v\n" +
			"\tLast Name: %v\n" +
			"\tEmail: %v\n" +
			"\tGender: %v\n" +
			"\tPublic: %v\n" +
			"\tJoin Date: %v\n" +
			"\tActive: %v\n" +
			"\t}\n\n",
			u.id, u.firstname, u.lastname,
			u.email, u.gender, u.public,
			u.joindate, u.active)
}

func LoginUserAccount(inputEmail string, inputPassword string) user {
	db := Database("login")
	query := fmt.Sprintf("SELECT * FROM user_account WHERE email='%s' AND password='%v';",
		inputEmail, Encrypt(inputPassword))
	r := db.QueryRow(query)
	defer db.Close()

	var (
		id        int
		age       int
		firstname string
		lastname  string
		email     string
		gender    string
		public    bool
		joindate  string
		active    bool
		password  string
	)

	err := r.Scan(&id, &age, &firstname, &lastname, &email, &gender, &public, &joindate, &active, &password)
	if err != nil {
		log.Fatal("Incorrect username or password.")
	}

	return user{
		id:        id,
		firstname: firstname,
		lastname:  lastname,
		email:     email,
		gender:    gender,
		public:    public,
		joindate:  joindate,
		active:    active,
	}
}

func DatabaseErrorHandler(err error) {
	e := fmt.Sprintf("%v", err)
	switch e {
	case "pq: duplicate key value violates unique constraint \"user_account_email_key\"":
		log.Fatal("User already exists.")
	}
}
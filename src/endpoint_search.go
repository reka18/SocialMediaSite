package main

import (
	"database/sql"
	json2 "encoding/json"
	"log"
	"net/http"
	"strings"
)

func searchGET(w http.ResponseWriter, r *http.Request) {

	CookieDebugger(r, "SEARCH ENDPOINT (GET)")

	username, ok := CompareTokens(w, r)
	if !ok {
		http.Redirect(w, r, "login", http.StatusSeeOther)
		return
	}

	RefreshCookie(username)

	db, _ := Database(DBNAME)
	defer db.Close()

	terms := ParseSearchQuery(r)
	users := SearchUser(terms, db)
	_, _ = w.Write(users)


}

func SearchHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	switch r.Method {
	case "GET":
		searchGET(w, r)
	case "POST":
		log.Println(Warn("No post method for search endpoint."))
	}

}

func SearchUser(input []string, db *sql.DB) []byte {

	limit := Min(len(input), 5)

	var idSet = make(map[int]bool)
	var userSet = make(map[int]SearchResult)

	for i := 0; i < limit; i++ {
		wildcard := "%" + input[i] + "%"
		r, e := db.Query("SELECT * FROM users WHERE username LIKE $1 OR firstname LIKE $1 OR lastname LIKE $1 OR email LIKE $1", wildcard)
		if e != nil {
			log.Println(Warn("Error making search query."))
		}

		var (
			ignore string // ignoring this encrypted password
			user   User
			result SearchResult
		)
		if r != nil {
			for r.Next() {
				e = r.Scan(&user.Id, &user.Age, &user.Firstname, &user.Lastname, &user.Email,
					&user.Username, &user.Public, &user.Joindate, &user.Active, &ignore, &user.Gender)
				if e != nil {
					log.Println(Warn("Error scanning user."))
				}

				if idSet[user.Id] {
					result.User = user
					result.Count = userSet[user.Id].Count + 1
					userSet[user.Id] = result
				} else {
					idSet[user.Id] = true
					result.User = user
					result.Count = 1
					userSet[user.Id] = result
				}
			}
		} else {
			log.Println(Warn("Query result is null."))
		}
	}

	var userResponse []SearchResult

	for _, v := range userSet {
		userResponse = append(userResponse, v)
	}

	json, e := json2.Marshal(userResponse)
	if e != nil {
		log.Println(Warn("Error making search query."))
	}
	log.Println(Info("Search result: ", string(json)))

	return json

}

func ParseSearchQuery(r *http.Request) []string {

	values, ok := r.URL.Query()["terms"]
	if !ok {
		log.Println(Warn("No search query terms specified."))
	} else {
		log.Println(Info("Found search terms: ", values))
	}
	value := values[0]

	return strings.Split(value, " ")

}
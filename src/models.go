package main

import (
	"net/http"
	"sync"
	"time"
)

type Config struct {
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type HTMLServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

type User struct {
	Id        int
	Username  string
	Firstname string
	Lastname  string
	Age       int
	Email     string
	Gender    string
	Public    bool
	Joindate  string
	Active    bool
}

type FollowedUser struct {
	User	User
	Mutual	bool
}

type SearchResult struct {
	User	User
	Count	int
}

type Post struct {
	Id				int
	UserId			int
	UserName		string // this is not stored in the post table
	Content			string
	UpVotes			int
	DownVotes		int
	Deleted			bool
	Date			time.Time
	FriendlyDate	string
}

type Message struct {
	Id				int
	FromId			int
	ToId			int
	Content			string
	Date			time.Time
	FriendlyDate	string
}

type Avatar struct {
	Id		int
	UserId	int
	Avatar	[]byte
}
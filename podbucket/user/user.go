package user

import (
	"net/http"
)

type User struct {
	isAdmin       bool
	isCreator     bool
	isContributor bool
	disabled      bool
	expire        string
	nick          string
	email         string
	name          string
	surname       string
	password      string
	age           int
}

func Handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		insertOne(res, req)
	}
}

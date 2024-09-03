package v1

import "net/http"

type Controller interface {
	AddUserHandler(w http.ResponseWriter, r *http.Request)
	GetMatchHandler(w http.ResponseWriter, r *http.Request)
}

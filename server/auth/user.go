package auth

import (
	"net/http"
)

func RegisterAuthControllers() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)
}

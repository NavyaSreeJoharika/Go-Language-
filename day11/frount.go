package day11

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/user", *uc)
	http.Handle("/user/", *uc)
}

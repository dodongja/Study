package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"goPractice/app/user/auth"
	"goPractice/app/user/signin"
	"goPractice/app/user/signup"
)

func MakeHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/signup", signup.SignUpHandler).Methods("POST")
	mux.HandleFunc("/signin", signin.SignInHandler).Methods("POST")

	//oauth
	mux.HandleFunc("/auth/google/login", auth.GoogleLoginHandler)
	mux.HandleFunc("/auth/google/callback", auth.GoogleAuthCallback)

	return mux
}

func main() {
	//cors도 따로 뺴야될듯?
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	http.ListenAndServe(":3000", cors.Handler(MakeHandler()))

}

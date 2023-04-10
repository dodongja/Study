package jwt

import (
	"fmt"
	"goPractice/app/user/signup"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func CreateAccessToken(w http.ResponseWriter, user *signup.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	accessUUID := uuid.NewString()
	// Create the JWT claims, which includes the username and expiry time
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = accessUUID
	atClaims["nickname"] = user.Nickname
	atClaims["exp"] = expirationTime

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// Create the JWT string

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE")))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return "", err
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   tokenString,
		Expires: expirationTime,
		// HttpOnly: true,
		//Secure:   true,
		//SameSite: http.SameSiteNoneMode,
	})

	fmt.Println("AccessToken", tokenString)

	return tokenString, nil

}

func CreateRefreshToken(w http.ResponseWriter) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	// Create the JWT claims, which includes the username and expiry time
	refreshUUID := uuid.NewString()
	// Create the JWT claims, which includes the username and expiry time
	atClaims := jwt.MapClaims{}
	atClaims["refresh_uuid"] = refreshUUID
	atClaims["exp"] = expirationTime

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	// Create the JWT string

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNATURE")))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return "", err
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "Refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
		// HttpOnly: true,
		//Secure:   true,
		//SameSite: http.SameSiteNoneMode,
	})

	fmt.Println("Refresh_token", tokenString)

	return tokenString, nil

}

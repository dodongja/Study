package jwt

import (
	"fmt"
	"goPractice/app/user/signup"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Nickname string `json:"nickname"`
	jwt.RegisteredClaims
}

func CreateAccessToken(w http.ResponseWriter, user *signup.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Nickname: user.Nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		Path:    "/",
		// HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	fmt.Println("AccessToken", tokenString)

	return tokenString, nil

}

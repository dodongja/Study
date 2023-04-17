package jwt

import (
	"fmt"
	"goPractice/app/user/signup"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func initializeRedisClient() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // 접근 url 및 port
		Password: "",               // password ""값은 없다는 뜻
		DB:       0,                // 기본 DB 사용
	})

	_, err := client.Ping().Result()

	return client, err
}

func executeSomething(uuid string, nickname string, exp time.Time) error {

	client, err := initializeRedisClient()

	if nil != err {
		panic(err)
	}

	now := time.Now()
	expiration := time.Unix(exp.Unix(), 0)

	saveErr := client.Set(uuid, nickname, expiration.Sub(now)).Err()
	if err != nil {
		return saveErr
	}

	return nil
}

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

	executeSomething(accessUUID, user.Nickname, expirationTime)

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

func CreateRefreshToken(w http.ResponseWriter, user *signup.User) (string, error) {
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

	executeSomething(refreshUUID, user.Nickname, expirationTime)

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

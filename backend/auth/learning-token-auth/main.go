package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	return token, err
}

type Claims struct {
	Id   string `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type Message struct {
	Message string
}

var secretKey = []byte("ini_adalah_secret")

func GenerateToken(id string, role string) (string, error) { //payload
	expirationTime := time.Now().Add(1 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Id:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // header sama payload yang di claim

	tokenString, err := token.SignedString(secretKey) //secret key
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		fmt.Println("error :", err)
		return "", err
	}

	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Query().Get("username")
	pass := r.URL.Query().Get("password")

	// logic pengecekan data di database, dimana, apakah username dan passwordnya ada

	if user == "" || pass == "" {
		OutputJSON(w, Message{"error bad request"})
		return
	}

	token, err := GenerateToken("001", "admin")
	if err != nil {
		OutputJSON(w, Message{"error generate token"})
		return
	}

	w.WriteHeader(200)
	OutputJSON(w, struct {
		Token string
	}{
		Token: token,
	})
}

func Secret(w http.ResponseWriter, r *http.Request) {
	// middleware, melakukan auth dari jwt token
	// header
	// query param
	tokenParam := r.URL.Query().Get("token")

	//sejenis verify untuk mendapatkan claim

	token, err := VerifyToken(tokenParam)

	if token != nil && err == nil {

		w.WriteHeader(200)
		OutputJSON(w, struct {
			Message string
		}{
			Message: "secret is : secret",
		})

	} else {
		w.WriteHeader(200)
		OutputJSON(w, struct {
			Message string
		}{
			Message: "error invalid token",
		})
	}

}

func main() {
	// token
	// header, payload, signature
	// claims

	http.HandleFunc("/login", Login)
	http.HandleFunc("/secret", Secret)

	fmt.Println("server running on port :8080")
	http.ListenAndServe(":8080", nil)
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	// struct to json
	data, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

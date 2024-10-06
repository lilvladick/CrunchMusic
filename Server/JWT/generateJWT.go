package main

// import (
// 	"fmt"
// 	"time"

// 	jwt "github.com/golang-jwt/jwt/v5"
// )

// func generateToken() {
// 	// Create a new token object, specifying signing method and the claims
// 	// you would like it to contain.
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"foo": "bar",
// 		"nbf": time.Now().Add(time.Hour * 72).Unix(),
// 	})

// 	// Sign and get the complete encoded token as a string using the secret
// 	tokenString, err := token.SignedString(hmacSampleSecret)

// 	fmt.Println(tokenString, err)

// }

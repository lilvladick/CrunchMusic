package main

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/golang-jwt/jwt/v5"
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

// func signToken() error {
// 	// get the token data from command line arguments
// 	tokData, err := loadData(*flagSign)
// 	if err != nil {
// 		return fmt.Errorf("couldn't read token: %w", err)
// 	}

// 	// parse the JSON of the claims
// 	var claims jwt.MapClaims
// 	if err := json.Unmarshal(tokData, &claims); err != nil {
// 		return fmt.Errorf("couldn't parse claims JSON: %w", err)
// 	}

// 	// add command line claims
// 	if len(flagClaims) > 0 {
// 		for k, v := range flagClaims {
// 			claims[k] = v
// 		}
// 	}

// 	// get the key
// 	var key interface{}
// 	key, err = loadData(*flagKey)
// 	if err != nil {
// 		return fmt.Errorf("couldn't read key: %w", err)
// 	}

// 	// get the signing alg
// 	alg := jwt.GetSigningMethod(*flagAlg)
// 	if alg == nil {
// 		return fmt.Errorf("couldn't find signing method: %v", *flagAlg)
// 	}

// 	// create a new token
// 	token := jwt.NewWithClaims(alg, claims)

// 	// sign the token
// 	out, err := token.SignedString(key)
// 	if err != nil {
// 		return fmt.Errorf("error signing token: %w", err)
// 	}
// 	fmt.Println(out)

// 	return nil
// }

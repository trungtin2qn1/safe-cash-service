package jwt

import (
	"errors"
	"log"
	"safe-cash-service/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//IssueToken ...
func IssueToken(id, email, storeID, role string, duration time.Duration) (string, error) {
	type ConsumerInfo struct {
		ID      string `json:"id,omitempty"`
		Email   string `json:"email,omitempty"`
		StoreID string `json:"store_id,omitempty"`
		Role    string `json:"role,omitempty"`
		jwt.StandardClaims
	}

	// time.Second * 86400: access Token
	// time.Second * 604800: refresh Token
	expire := time.Now().Add(duration).Unix()

	log.Println("[issue-token] storeID:", storeID)

	// Create the Claims
	claims := &ConsumerInfo{
		id,
		email,
		storeID,
		role,
		jwt.StandardClaims{
			Issuer:    "Tin",
			ExpiresAt: expire,
		},
	}
	return SignToken(claims)
}

//SignToken sign claims
func SignToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(signKey)

	if nil != err {
		log.Println("Error while signing the token")
		log.Printf("Error signing token: %v\n", err)
		go utils.LogErrToFile(err.Error())
		return ss, err
	}

	return ss, nil
}

//VerificationToken ...
func VerificationToken(tokenString string) (string, string, string, string, error) {
	type ConsumerInfo struct {
		ID      string `json:"id,omitempty"`
		Email   string `json:"email,omitempty"`
		StoreID string `json:"store_id,omitempty"`
		Role    string `json:"role,omitempty"`
		jwt.StandardClaims
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &ConsumerInfo{},
		func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			return verifyKey, nil
		})

	if err != nil {
		go utils.LogErrToFile(err.Error())
		return "", "", "", "", err
	}
	claims, ok := token.Claims.(*ConsumerInfo)

	if !ok {
		return "", "", "", "", errors.New("invalid token")
	}
	return claims.ID, claims.Email, claims.StoreID, claims.Role, err
}

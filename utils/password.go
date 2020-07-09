package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

//Generate a salted hash for the input string
func Generate(s string) (string, error) {
	saltedBytes := []byte(s + getPepper())
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		// go LogErrToFile(err.Error())
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare ...
func Compare(hash string, s string) (bool, error) {
	incoming := []byte(s + getPepper())
	existing := []byte(hash)

	err := bcrypt.CompareHashAndPassword(existing, incoming)
	if err != nil {
		log.Println("err 0:", err)
		// go LogErrToFile(err.Error())
		return false, err
	}
	return true, nil
}

func getPepper() string {
	return "pepper"
}

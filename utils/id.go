package utils

import "github.com/google/uuid"

//GenerateUUID ...
func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

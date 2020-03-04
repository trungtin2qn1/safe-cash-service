package main

import (
	"safe-cash-service/db"
	"safe-cash-service/routers"
	"safe-cash-service/utils/jwt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	// Connect to database
	db.Init(nil)
	defer db.CloseDB()

	// Load keys for jwt
	jwt.LoadRSAKeys()

	// Setup routers ....
	routers.SetUpRouter()

}

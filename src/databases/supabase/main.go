package supabase

import (
	"fmt"
	"project/src/helpers"

	"github.com/go-pg/pg/v10"
)

var DB pg.DB
var Err error

func Connect() {
	addr, _ := helpers.EnvGetProperty("SUPABASE_ADDR")
	user, _ := helpers.EnvGetProperty("SUPABASE_USER")
	dbName, _ := helpers.EnvGetProperty("SUPABASE_DB")
	password, _ := helpers.EnvGetProperty("SUPABASE_PASSWORD")

	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     user,
		Password: password,
		Database: dbName,
	})

	DB = *db

	fmt.Println("Database is connected")
}

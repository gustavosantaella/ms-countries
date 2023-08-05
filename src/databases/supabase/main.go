package supabase

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

var DB pg.DB
var Err error

func Connect() {
	db := pg.Connect(&pg.Options{
		Addr:     "db.zgrnrageaclbfipotihp.supabase.co:5432",
		User:     "postgres",
		Password: "q?5bN!DX6PSHEwA",
		Database: "postgres",
	})

	DB = *db

	fmt.Println("Database is connected")
}

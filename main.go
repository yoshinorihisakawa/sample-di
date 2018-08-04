package main

import (
	"database/sql"
	"github.com/yoshinorihisakawa/sample-di/repository"
	"github.com/yoshinorihisakawa/sample-di/service"
)

func main() {
	db, err := sql.Open("mysql", "root:@/database")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)

	// usを使ってゴニョゴニョ
}

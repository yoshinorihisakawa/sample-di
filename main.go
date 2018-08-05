package main

import (
	"github.com/yoshinorihisakawa/sample-di/database"
	"github.com/yoshinorihisakawa/sample-di/repository"
	"github.com/yoshinorihisakawa/sample-di/service"
)

func main() {
	d := database.NewDatabase("mysql", "root:@/database")
	ur := repository.NewUserRepository(&d)
	us := service.NewUserService(ur)

	// usを使ってゴニョゴニョ
}

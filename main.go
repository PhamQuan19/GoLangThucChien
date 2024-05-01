package main

import (
	"my-app/db"
	"my-app/handler"
	repoimpl "my-app/repository/repo_impl"
	"my-app/router"

	// "my-app/handler"
	// "net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	sql := &db.Sql{
		Host: "localhost",
		Port: 5432,
		UserName: "postgres",
		Password: "111999",
		DbName: "golang",

	}

	sql.Connect()
	defer sql.Close()

	e := echo.New()
	// e.GET("/", handler.Welcome)
	// e.GET("/user/sign-in",handler.HandleSignIn)
	// e.GET("/user/sign-up",handler.HandleSignUp)
	userHanler := handler.UserHandler{
		UserRepo: repoimpl.NewUserRepo(sql),
	}

	api:=router.API{
		Echo: e,
		UserHandler: userHanler,
	}

	api.SetupRouter()
	
	e.Logger.Fatal(e.Start(":1323"))
}




package main

import (
	"log"

	g "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/leeleeleeee/web-app/auth"
	"github.com/leeleeleeee/web-app/conf"
	c "github.com/leeleeleeee/web-app/controller"
	m "github.com/leeleeleeee/web-app/model"
)

func Logger() g.HandlerFunc {
	return func(c *g.Context) {
		c.Next()
		// Set example variable
		c.Set("example", "12345")

		// before request
		// c.Error(errors.New("Something Error"))
		// err := c.Errors.Last()
		// if err == nil {

		// }

		// after request

		log.Print("logger1 - func1")

		// access the status we are sending
		status := c.Writer.Status()
		log.Print("logger1 - func1", status)
	}
}

func main() {
	err := godotenv.Load()
	pg := &m.Postgres{}
	if err != nil {
		panic("Please create a .env file")
	}
	conf.RedisInit()

	pg.ConnectDatabse()
	pg.MigrateDatabase()
	m.Gdb = pg.GetDB()
	router := g.Default()
	router.Use(Logger())

	defer func() {

	}()

	authorized := router.Group("/api/auth")
	c.DoInit(authorized, auth.AuthController{})

	api := router.Group("/api/v1", auth.CheckJwt())
	c.DoInit(api, c.UserController{})

	router.Run(":8080")
}

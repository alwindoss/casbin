package main

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"

	casbinmw "github.com/alwindoss/casbin"
)

func main() {
	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		panic(err)
	}
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		// handle err
	}
	if ok == true {
		// permit alice to read data1
	} else {
		// deny the request, show an error
	}

	app := fiber.New()

	app.Use(casbinmw.New(casbinmw.Config{
		Enforcer: e,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}

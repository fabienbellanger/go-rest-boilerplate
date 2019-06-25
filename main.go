package main

import "github.com/fabienbellanger/go-rest-boilerplate/commands"

func main() {
	// Lancement de Cobra
	commands.Execute()
}

// package main

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"strconv"

// 	"github.com/labstack/echo"
// )

// type User struct {
// 	Username  string
// 	Password  string
// 	Lastname  string
// 	Firstname string
// }

// func main() {
// 	e := echo.New()

// 	e.GET("/", func(c echo.Context) error {
// 		// var user User
// 		// users := make([]User, 0)

// 		// for i := 0; i < 100000; i++ {
// 		// 	user = User{
// 		// 		Username:  "ffgfgfghhfghfhgfgfhgfghfghfhgfhgfh" + strconv.Itoa(i),
// 		// 		Password:  "gjgjghjgjhgjhghjfrserhkhjhklljjkbhjvftxersgdghjjkhkljkbhftd",
// 		// 		Lastname:  "njuftydfhgjkjlkjlkjlkhjkhu",
// 		// 		Firstname: "jkggkjkl,,lm,kljkvgf"}

// 		// 	users = append(users, user)
// 		// }

// 		// c.Response().WriteHeader(http.StatusOK)
// 		// for _, user := range users {
// 		// 	if err := json.NewEncoder(c.Response()).Encode(user); err != nil {
// 		// 		return err
// 		// 	}
// 		// 	c.Response().Flush()
// 		// }
// 		// return nil

// 		resp := c.Response()
// 		resp.WriteHeader(http.StatusOK)

// 		if _, err := io.WriteString(resp, "["); err != nil {
// 			return err
// 		}
// 		enc := json.NewEncoder(resp)
// 		for i := 0; i < 100000; i++ {
// 			if i > 0 {
// 				if _, err := io.WriteString(resp, ","); err != nil {
// 					return err
// 				}
// 			}
// 			user := User{
// 				Username:  "ffgfgfghhfghfhgfgfhgfghfghfhgfhgfh" + strconv.Itoa(i),
// 				Password:  "gjgjghjgjhgjhghjfrserhkhjhklljjkbhjvftxersgdghjjkhkljkbhft",
// 				Lastname:  "njuftydfhgjkjlkjlkjlkhjkhu",
// 				Firstname: "jkggkjkl,,lm,kljkvgf",
// 			}
// 			if err := enc.Encode(user); err != nil {
// 				return err
// 			}
// 		}
// 		if _, err := io.WriteString(resp, "]"); err != nil {
// 			return err
// 		}

// 		return nil
// 	})

// 	e.Logger.Fatal(e.Start(":1323"))
// }

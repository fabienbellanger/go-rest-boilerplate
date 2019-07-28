package main

import "github.com/fabienbellanger/go-rest-boilerplate/commands"

func main() {
	// Lancement de Cobra
	commands.Execute()
}

// User type
// type User struct {
// 	ID        uint64
// 	Username  string
// 	Password  string
// 	Lastname  string
// 	Firstname string
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt time.Time
// }

// func main() {
// 	users := benchmarkEcho()
// 	// fmt.Println(users)

// 	time.Sleep(20 * time.Second)

// 	s := make([]User, 0, 0)
// 	copy(s, users)
// 	users = nil

// 	time.Sleep(30 * time.Second)
// }

// func benchmarkEcho() []User {
// 	var user User
// 	var i uint64
// 	var n uint64 = 100000

// 	users := make([]User, n, n)

// 	for i = 0; i < n; i++ {
// 		user.ID = i + 1
// 		user.Username = "Username jjhfjhfjjfhjf jdhjfhdjhfjd"
// 		user.Lastname = "Lastname jkfdjkfjkdjkfdkfkd"
// 		user.Firstname = "Firstname jjhfjhfjjfhjf jdhjfhdjhfjd"
// 		user.Password = "SFSFDSdddfjdskjfkjdk345345GCTHG5ER9?TG9°?4TN34°AT8°34NTCNRT92°C5T2NV5°E9RNTG°5E9NCG9°ENR9NZ°ERNG9REZNG°ERZNG°RNGR"
// 		user.CreatedAt = time.Now()
// 		user.UpdatedAt = time.Now()
// 		user.DeletedAt = time.Now()

// 		users[i] = user
// 	}
// 	// fmt.Println("Before first copy")
// 	// time.Sleep(20 * time.Second)
// 	// usersLight := make([]User, 20000, 20000)
// 	// copy(usersLight, users)
// 	// users = nil

// 	// time.Sleep(20 * time.Second)
// 	// fmt.Println("After first copy")

// 	// return usersLight
// 	return users
// }

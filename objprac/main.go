package main

import (
    "fmt"
)

type User struct {
    ID    int
    Name  string
    Email string
}

var users []User
var idCounter = 1

// CREATE
func createUser(name, email string) User {
    user := User{ID: idCounter, Name: name, Email: email}
    users = append(users, user)
    idCounter++
    return user
}

// READ
func getUser(id int) *User {
    for _, user := range users {
        if user.ID == id {
            return &user
        }
    }
    return nil
}

// UPDATE
func updateUser(id int, name, email string) bool {
    for i := range users {
        if users[i].ID == id {
            users[i].Name = name
            users[i].Email = email
            return true
        }
    }
    return false
}

// DELETE
func deleteUser(id int) bool {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return true
        }
    }
    return false
}

func main() {
    u1 := createUser("Alice", "alice@example.com")
    u2 := createUser("Bob", "bob@example.com")

    fmt.Println("Users after creation:", users)

    user := getUser(u1.ID)
    if user != nil {
        fmt.Println("Retrieved:", *user)
    }

    updated := updateUser(u2.ID, "Bobby", "bobby@example.com")
    fmt.Println("Update success?", updated)
    fmt.Println("Users after update:", users)

    deleted := deleteUser(u1.ID)
    fmt.Println("Delete success?", deleted)
    fmt.Println("Users after deletion:", users)
}

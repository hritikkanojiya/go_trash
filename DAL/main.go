package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ConnectionConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func InitAppMeta() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func SetupPGConfig() ConnectionConfig {
	return ConnectionConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
}

func main() {
	InitAppMeta()
	config := SetupPGConfig()
	poolConnection, err := GetPGConnection(config)
	if err != nil {
		panic(err)
	}
	defer poolConnection.Close()

	fmt.Println("Successfully connected to the database!")

	userRepo := NewUserRepository(poolConnection)

	fmt.Println("Read Record")
	users, err := userRepo.GetAllUsers()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	fmt.Println("Create Record")

	GMUser := User{
		Name:  "Genetic Minds",
		Email: "we@geneticminds.com",
	}

	err = userRepo.CreateUser(GMUser)
	if err != nil {
		panic(err)
	}

	fmt.Println("User Created")

	fmt.Println("Read Again")

	newUsers, err := userRepo.GetAllUsers()
	if err != nil {
		panic(err)
	}

	for _, user := range newUsers {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	latestUser, err := userRepo.GetLatestUser()
	if err != nil {
		panic(err)
	}

	userID := latestUser.ID

	user, err := userRepo.GetUserByID(*userID)
	if err != nil {
		panic(err)
	}

	user.Name = "GM"
	user.Email = "team@geneticminds.com"

	fmt.Println("Update Record")

	err = userRepo.UpdateUser(user)
	if err != nil {
		panic(err)
	}

	fmt.Println("User Updated")

	fmt.Println("Read Again")

	updateUsers, err := userRepo.GetAllUsers()
	if err != nil {
		panic(err)
	}

	for _, user := range updateUsers {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

	fmt.Println("Delete Record")

	err = userRepo.DeleteUser(*userID)
	if err != nil {
		panic(err)
	}

	fmt.Println("User Deleted")

	fmt.Println("Read Again")

	remainingUsers, err := userRepo.GetAllUsers()
	if err != nil {
		panic(err)
	}

	for _, user := range remainingUsers {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
	}

}

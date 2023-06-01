package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
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

func executeFunction(pool *pgxpool.Pool, functionName string, args ...interface{}) (pgx.Row, error) {
	query := fmt.Sprintf("SELECT * FROM %s(%s)", functionName, generateArgumentPlaceholders(len(args)))
	return pool.QueryRow(context.Background(), query, args...), nil
}
func generateArgumentPlaceholders(count int) string {
	placeholders := ""
	for i := 1; i <= count; i++ {
		if i > 1 {
			placeholders += ", "
		}
		placeholders += fmt.Sprintf("$%d", i)
	}
	return placeholders
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

	rows, err := executeFunction(poolConnection, "get_all_users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ID:", id, "Name:", name, "Email:", email)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	/*
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
	*/

}

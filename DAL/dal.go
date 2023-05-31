package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	ID    *int
	Name  string
	Email string
}

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}

func (ur *UserRepository) CreateUser(user User) error {
	_, err := ur.pool.Exec(context.Background(), "INSERT INTO users(name, email) VALUES($1, $2)", user.Name, user.Email)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) GetUserByID(id int) (*User, error) {
	row := ur.pool.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", id)
	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetAllUsers() ([]User, error) {
	rows, err := ur.pool.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) GetLatestUser() (*User, error) {
	var user User

	err := ur.pool.QueryRow(context.Background(), "SELECT id, name, email FROM users ORDER BY id DESC LIMIT 1").
		Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateUser(user *User) error {
	_, err := ur.pool.Exec(context.Background(), "UPDATE users SET name = $1, email = $2 WHERE id = $3", user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	_, err := ur.pool.Exec(context.Background(), "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

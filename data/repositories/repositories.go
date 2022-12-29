package repositories

import (
	connections "connections"
	entities "entities"
	"fmt"
	"time"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func SchemaMigrations() {
	fmt.Println("repositories Schema Migrations")
	db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Migrate the schema
	db.Migrator().CreateTable(&entities.User{})
}

func CreateNewUser(user entities.User, password string, is_email_verified bool) entities.User {
	fmt.Println("repositories Create New User")

	db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.Exec("INSERT INTO users (created_at, email, password) VALUES (?,?,?)", time.Now(), user.Email, password)

	return user
}

func UpdateUser(user entities.User) {
	fmt.Println("repositories Update User")
	db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")

	}
	db.Exec("UPDATE users SET password = ? WHERE email = ?", user.Password, user.Email)
}

func GetAllUsers() []entities.User {
	fmt.Println("repositories Get All Users")
	db, err := gorm.Open(sqlserver.Open(connections.ConnectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")

	}
	var users []entities.User
	db.Raw("SELECT * FROM users").Scan(&users)
	return users
}

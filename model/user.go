package model

import (
    "database/sql"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/go-sql-driver/mysql"
	
)

var db *sql.DB

type User struct {
	ID       int
	Username string
	Password string
}

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
}
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsersByID(id int) (*User, error) {
	var user User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers() ([]User, error) {
	var users []User
	rows, err := db.Query("SELECT id, username, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func SaveUserToDB(user User) error {
	// 插入用戶資料
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

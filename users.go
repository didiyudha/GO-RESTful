package main

import (
	"errors"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID        int64  `json:"Id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  []byte `json:"-"`
}

// NewUser retrun a pointer to a User
func NewUser() *User {
	return new(User)
}

// GetAllUsers return slice of User that contains all users in database
func (u *User) GetAllUsers() ([]User, error) {
	users := make([]User, 0)
	db, err := GetDB()
	if err != nil {
		return users, err
	}
	defer db.Close()
	rows, err := db.Query(
		`SELECT id,
      username,
      first_name,
      last_name
     FROM users;
    `)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		usr := User{}
		rows.Scan(&usr.ID, &usr.Username, &usr.FirstName, &usr.LastName)
		users = append(users, usr)
	}
	log.Println("Users: ", users)
	return users, nil
}

// Save user
func (u *User) Save() error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Close()
	bPass, err := getDefaultPassword()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		`INSERT INTO users (username, first_name, last_name, password)
     VALUES (?, ?, ?, ?)`,
		u.Username, u.FirstName, u.LastName, bPass)
	if err != nil {
		return err
	}
	return nil
}

// FindByID return user by ID
func (u *User) FindByID(ID string) (User, error) {
	if ID == "" {
		return User{}, errors.New("ID can not be empty")
	}
	_, err := strconv.Atoi(ID)
	if err != nil {
		return User{}, errors.New("ID must be a number")
	}
	db, err := GetDB()
	usr := User{}
	if err != nil {
		return usr, err
	}
	defer db.Close()
	rows, err := db.Query(
		`SELECT id,
      username,
      first_name,
      last_name
    FROM users
    WHERE id = ` + ID)
	if err != nil {
		return usr, err
	}
	if rows.Next() {
		rows.Scan(&usr.ID, &usr.Username, &usr.FirstName, &usr.LastName)
		return usr, nil
	}
	return User{}, errors.New("No data found")
}

// Update existing user
func (u *User) Update(newUsr User) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(
		`UPDATE users SET username = ?, first_name = ?, last_name=? WHERE id = ?`,
		newUsr.Username, newUsr.FirstName, newUsr.LastName, newUsr.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete user from database
func (u *User) Delete(ID string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM users WHERE id = ?`, ID)
	return err
}

func getDefaultPassword() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)
}

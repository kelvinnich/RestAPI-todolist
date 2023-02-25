package repository

import (
	"authenctications/model"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UsersRepository interface {
	CreateUsers(m model.Users) (*model.Users, error)
	UpdateUsers(id string, m model.Users) (*model.Users, error)
	IsDuplicateEmail(email string) (bool, error) 
	VerifyCredential(email string, password string) interface{}
	FindUserById(id string) (*model.Users, error)
	Profile(id string) (model.Users, error)
}

type usersConnection struct {
	db *sql.DB
}


// This function creates a new user in the users table of the database. It takes in a model.Users struct as an argument and hashes the password before inserting it into the database. It then executes an INSERT query with the given values and returns a pointer to the model.Users struct if successful, or an error if not.
func (db *usersConnection) CreateUsers(m model.Users) (*model.Users, error) {
	m.Password = hashedPassword([]byte(m.Password))
	query := `INSERT INTO users(id, username, email, password, address) VALUES($1, $2, $3, $4, $5)`

	_, err := db.db.Exec(query, m.ID, m.Username, m.Email, m.Password, m.Address)
	if err != nil {
		return nil, err
	}

	return &m, nil
}


// This function updates a user in the users table of a database. It takes an id string and a model.Users struct as parameters. It creates a query to update the username, email, password, and address fields of the user with the given id. If successful, it returns a pointer to the model.Users struct and nil for error; otherwise, it returns nil for model.Users and an error.
func (db *usersConnection) UpdateUsers(id string, m model.Users) (*model.Users, error) {
	query := `UPDATE users SET username = $1, email = $2, password = $3, address = $4 WHERE id = $5`

	_, err := db.db.Exec(query, m.Username, m.Email, m.Password, m.Address, id)
	if err != nil {
		return nil, err
	}

	return &m, nil
}


// This function checks if a given email is already registered in the users table of the usersConnection database. It takes in an email string as a parameter and returns a boolean and an error. The query checks to see if there is an existing user with the same email address, and if there is, it returns false and no error. If there isn't, it returns true and no error. If there is an error, it prints out a log message and returns false and nil.
func (db *usersConnection) IsDuplicateEmail(email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 LIMIT 1)`
	var exists bool
	err := db.db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		log.Printf("failed to register user repository %v", err)
		return false,nil
	}

	return !exists, nil
}


// This function is used to verify a user's credentials. It takes two strings, an email and a password, as parameters. It then queries the database for the user associated with the given email address and stores the result in a variable of type model.Users. If no user is found, it logs an error message and returns nil. Otherwise, it compares the given password with the stored password using bcrypt.CompareHashAndPassword(). If the passwords do not match, it logs an error message and returns nil. Otherwise, it returns the user data from model.Users.
func (db *usersConnection) VerifyCredential(email string, password string) interface{} {
	var user model.Users
	query := `SELECT id, username, email, password, address FROM users WHERE email = $1 LIMIT 1`
	err := db.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No user found with that email")
			return nil
		} else {
			log.Println("Error getting user data from database")
			return nil
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Password does not match")
		return nil
	}

	return user
}


// This function takes in a string parameter (id) and uses it to query the users table in the database. It then scans the result of the query and stores it in a model.Users struct. If an error is encountered, it returns nil and the error, otherwise it returns a pointer to the model.Users struct and nil for the error.
func (db *usersConnection) FindUserById(id string) (*model.Users, error) {
	var user model.Users
	query := `SELECT * FROM users WHERE id = $1`

	err := db.db.QueryRow(query, id).Scan(&user.Username, &user.Email, &user.Password, &user.Address, &user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &user, nil
}


// This function Profile() takes in a string id as an argument and returns a model.Users object and an error. It queries the users table in the database for the row with the given id and scans it into a user object. If there is an error, it will return an empty user object and nil error. If there is no row with the given id, it will return an empty user object and nil error. Otherwise, it will return the user object populated with data from the database and nil error.
func (db *usersConnection) Profile(id string) (model.Users, error) {
	var user model.Users
	query := `SELECT id,username,email,password,address FROM users WHERE id = $1`

	row := db.db.QueryRow(query, id)
	err := row.Scan(&user.ID,&user.Username, &user.Email, &user.Password, &user.Address)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Users{}, nil
		}
		return model.Users{}, nil
	}

	return user, nil
}

func hashedPassword(pw []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password %v", err)
	}

	return string(hash)
}

func NewUserRepository(db *sql.DB) UsersRepository {
	return &usersConnection{
		db: db,
	}
}

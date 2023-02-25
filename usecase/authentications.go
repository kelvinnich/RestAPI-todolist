package usecase

import (
	"authenctications/dto"
	"authenctications/model"
	"authenctications/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type Authentications interface {
	CreateUsers(userDto dto.RegisterUsersDTO) *model.Users
	IsDuplicateEmail(email string) (bool)
	VerifyCredential(email string, password string) interface{}
	FindUserById(id string) (*model.Users, error)
}

type authenctications struct {
	userRepository repository.UsersRepository
}


// This function creates a user using the authentication service. It takes in a RegisterUsersDTO as an argument and uses it to fill a model.Users struct. It then calls the userRepository.CreateUsers method to create the user, and returns the created user if successful. If any errors occur, they are logged.
func(a *authenctications)CreateUsers(userDto dto.RegisterUsersDTO) *model.Users{
	var user model.Users
	err := smapping.FillStruct(&user, smapping.MapFields(&userDto))
	if err != nil {
		log.Printf("failed to map %v", err)
	}

	u,err := a.userRepository.CreateUsers(user)
	if err != nil {
		log.Printf("failed to create user usecase %v", err)
	}

	return u
}


// This function checks if a given email is a duplicate in the user repository. It takes an email string as an argument and returns a boolean value. If there is an error, it prints out a log message with the error.
func (a *authenctications) IsDuplicateEmail(email string) bool {
	result,err := a.userRepository.IsDuplicateEmail(email)
	if err != nil {
		log.Printf("failed to register user, email is duplicate %v", err)
	}

	return result
}


// This function is part of the authenctications struct and is used to verify credentials. It takes two strings, an email and a password, as parameters. It uses the userRepository to verify the credentials and then checks if the length of the password is less than bcrypt's minimum cost. If it is, it returns nil. If not, it compares the hash of the password with the given password using bcrypt and returns nil if they don't match. Otherwise, it returns v which is a model.Users type.
func (a *authenctications) VerifyCredential(email string, password string) interface{}{
	verify := a.userRepository.VerifyCredential(email, password)
	if verify == nil {
			return nil
	}
	v := verify.(model.Users)
	if len(v.Password) < bcrypt.MinCost{
		return nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(password))
	if err != nil {
			log.Print("Invalid credentials", err)
			return nil
	}
	return v
}



//This function takes in a pointer to an authentication object and a string representing an ID as parameters. It then calls the FindUserById method on the userRepository property of the authentication object, passing in the ID as an argument. The function returns a pointer to a model.Users object and an error value.
func(a *authenctications)FindUserById(id string) (*model.Users, error){
	return a.userRepository.FindUserById(id)
}

func NewAuthentication(userrepo repository.UsersRepository) Authentications{
	return &authenctications{
		userRepository: userrepo,
	}
}




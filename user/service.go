package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
	GetUserById(ID int) (User, error)
}

type service struct {
	repositry Repository
}

func NewService(repositry Repository) *service {
	return &service{repositry}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {

	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repositry.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repositry.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repositry.FindByEmail(email)
	if err != nil {
		return false, nil
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(ID int, filelocation string) (User, error) {
	//dapatkan user berdasarkan ID
	//update atribut avatar file name
	//simpan perubahan avatar file name

	user, err := s.repositry.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = filelocation

	updateUser, err := s.repositry.Update(user)
	if err != nil {
		return updateUser, err
	}

	return updateUser, nil

}

func (s *service) GetUserById(ID int) (User, error) {
	user, err := s.repositry.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with thath ID")
	}

	return user, nil

}

// mapping struct input ke struct user
// simpan struct User melalui repository

package user_service

import (
	"errors"
	"strings"

	models "github.com/NOTMKW/API/internal/model"
	repository "github.com/NOTMKW/API/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	user := &models.User{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     strings.ToLower(req.Email),
		Password:  string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		if strings.Contains(err.Error(), "unique constraint failed") ||
			strings.Contains(err.Error(), "duplicate entry") ||
			strings.Contains(err.Error(), "duplicate key value") {
			return nil, errors.New("user with this email already exists")
		}
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

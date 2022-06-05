package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hberkayozdemir/hibemi-be/helpers"
	"github.com/hberkayozdemir/hibemi-be/internal/client"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Service struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		Repository: repository,
	}
}

func (s *Service) RegisterUser(userDTO UserDTO) (*User, error) {
	registeredUser, err := s.Repository.GetUserByEmail(userDTO.Email)
	if registeredUser != nil {
		return nil, UserAlreadyExistError
	}

	hashedPassword, _ := helpers.HashPassword(userDTO.Password)

	user := User{
		ID:              helpers.GenerateUUID(8),
		FirstName:       userDTO.FirstName,
		LastName:        userDTO.LastName,
		Email:           userDTO.Email,
		Phone:           userDTO.Phone,
		Password:        hashedPassword,
		IsEmailActivate: false,
		UserType:        "user",
	}

	newUser, err := s.Repository.RegisterUser(user)

	if err != nil {
		return nil, err
	}

	activationCode := strconv.Itoa(rand.Intn(999999-100000) + 100000)

	err = s.Repository.AddActivationCode(user.Email, activationCode)
	if err != nil {
		return nil, err
	}

	err = client.SendMail(user.Email, "Hesabını Doğrula", "Hesabınızı doğrulamak için lütfen kodu giriniz: "+activationCode)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (s *Service) LoginUser(userCredentialsDTO UserCredentialsDTO) (*Token, *fiber.Cookie, error) {
	user, err := s.Repository.GetUserByEmail(userCredentialsDTO.Email)
	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}
	if !helpers.CheckPasswordHash(userCredentialsDTO.Password, user.Password) || !user.IsEmailActivate {
		return nil, nil, err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserType: user.UserType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    user.ID,
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return nil, nil, err
	}

	cookie := fiber.Cookie{
		Name:    "user-token",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}

	return &Token{
		Token: token,
	}, &cookie, nil
}

func (s *Service) DeleteUser(userID string) error {
	err := s.Repository.DeleteUser(userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) ActivateUser(email, code string) (*User, error) {
	registeredUser, err := s.Repository.GetUserByEmail(email)
	if err != nil {
		return nil, UserNotFound
	}

	if registeredUser.IsEmailActivate {
		return nil, UserAlreadyActivated
	}

	err = s.Repository.DeleteActivationCode(code)
	if err != nil {
		return nil, err
	}

	user, err := s.Repository.ActivateUser(registeredUser.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUsersPageableResponse(pageNumber, size int) (*UsersPageableResponse, error) {
	users, totalElements, err := s.Repository.GetUsers(pageNumber, size)
	if err != nil {
		return nil, err
	}
	page := Page{
		Number:        pageNumber,
		Size:          size,
		TotalElements: totalElements,
		TotalPages:    int(math.Ceil(float64(totalElements) / float64(size))),
	}

	return &UsersPageableResponse{
		Users: users,
		Page:  page,
	}, nil
}

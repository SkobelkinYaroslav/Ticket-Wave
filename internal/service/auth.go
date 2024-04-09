package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"ticket_wave/internal/models"
	"time"
)

func (s Service) RegisterService(req models.Participant) error {
	fmt.Println("register service hit")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	fmt.Println("hashed password: ", string(hashedPassword))

	req.Password = string(hashedPassword)

	if err := s.AuthRepository.CreateUserRepo(req); err != nil {
		return err
	}

	fmt.Println("service: user created")

	return nil
}

func (s Service) LoginService(req models.Participant) (string, error) {
	user, err := s.AuthRepository.GetUserRepo(req)

	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", err
	}

	token, err := s.createJWTService(user)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s Service) CheckTokenService(tokenString string) (models.Participant, error) {
	if tokenString == "" {
		return models.Participant{}, fmt.Errorf("empty token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return models.Participant{}, fmt.Errorf("error while parsing token")
	}

	var user models.Participant

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return models.Participant{}, fmt.Errorf("token expired")
		}

		req := models.Participant{
			Email: claims["sub"].(string),
		}
		user, err = s.AuthRepository.GetUserRepo(req)

		if err != nil {
			return models.Participant{}, fmt.Errorf("error while getting user")
		}
	} else {
		return models.Participant{}, fmt.Errorf("error while getting claims")
	}

	return user, nil
}

func (s Service) createJWTService(user models.Participant) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

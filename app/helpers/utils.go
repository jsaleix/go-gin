package helpers

import (
	"api/config"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var Validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	log.Printf("a %s -- b %s, res %t", userPassword, providedPassword, err == nil)
	return err == nil
}

type SignDetails struct {
	Email     string
	Uid       string
	User_Type string
	jwt.StandardClaims
}

func GenerateAllTokens(email string, userType string, uid string) (token string, refreshToken string, err error) {
	claims := &SignDetails{
		Email:     email,
		Uid:       uid,
		User_Type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, tokenErr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.SECRET))
	refreshToken, refreshErr := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(config.SECRET))

	if tokenErr != nil || refreshErr != nil {
		return
	}

	return token, refreshToken, err
}

func ValidateToken(signedToken string) (claims *SignDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(config.SECRET), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignDetails)

	if !ok {
		msg = "Token is invalid"
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Token has expired"
		return
	}

	return claims, msg

}

package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(id uint) (string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file, %s\n", err)
		return "", err
	}

	token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	if err != nil {
		return "", err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sup": id,
		"exp": time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TokenValid(c *gin.Context) error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file, %s\n", err)
		return err
	}

	tokenString, ok := ExtractToken(c)

	if !ok {
		return errors.New("there is no token in header")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ExtractToken(c *gin.Context) (string, bool) {
	token, err := c.Cookie("Authorization")

	if err != nil {
		return "", false
	}

	return token, true
}

func ExtractTokenID(c *gin.Context) (int, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("Error loading .env file, %s\n", err)
		return 0, err
	}

	tokenString, ok := ExtractToken(c)
	if !ok {
		return 0, errors.New("there is no token in your header")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return 0, fmt.Errorf("can't fetch id from claims")

	}
	fmt.Println(claims)
	value, ok := claims["sub"].(float64)
	if !ok {
		// Handle non-float64 value
	}
	uid, err := strconv.ParseUint(fmt.Sprintf("%d", int64(value)), 10, 64)
	if err != nil {
		return 0, err
	}
	return int(uid), nil
}

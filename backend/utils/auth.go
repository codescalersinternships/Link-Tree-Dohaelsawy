package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrNoTokenCookie        = errors.New("there is no token in your cookies, login in first")
	ErrInvalidToken         = errors.New("invalid token")
	ErrWrongTokenMethod     = errors.New("unexpected signing method")
	ErrExtractIDTokenClaims = errors.New("can't fetch id from claims")
)

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CreateToken(id uint, token_lifespan int, secretToken string) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sup": id,
		"exp": time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString([]byte(secretToken))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func TokenValid(c *gin.Context) error {

	config, err := NewConfigController()
	if err != nil {
		return err
	}

	tokenString, ok := ExtractToken(c)
	if !ok {
		return ErrNoTokenCookie
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return err
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

	config, err := NewConfigController()
	if err != nil {
		return 0, err
	}

	tokenString, ok := ExtractToken(c)
	if !ok {
		return 0, ErrNoTokenCookie
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrWrongTokenMethod
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return 0, ErrExtractIDTokenClaims
	}

	value, ok := claims["sup"].(float64)
	if !ok {
		return 0, ErrExtractIDTokenClaims
	}

	uid, err := strconv.ParseUint(fmt.Sprintf("%d", int64(value)), 10, 64)
	if err != nil {
		return 0, err
	}

	return int(uid), nil
}

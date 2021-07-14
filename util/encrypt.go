package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strconv"
	"time"
)

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomPassword(n int) string {
	var letters = []rune(fmt.Sprintf("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789%s", strconv.Itoa(int(time.Now().Unix()))))
	password := make([]rune, n)
	for i := range password {
		password[i] = letters[rand.Intn(len(letters))]
	}
	return string(password)
}

type JwtClaim struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func GenerateToken(username string) string {
	claim := JwtClaim{Username: username}
	secret := Config.GetString("jwt.secret")
	claim.IssuedAt = time.Now().Unix()
	claim.ExpiresAt = time.Now().Add(time.Second * time.Duration(100)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		print("生成jwt token失败")
		print(err.Error())
		panic(err)
	}
	return signedToken
}

func CheckToken(token string) (*JwtClaim, bool) {
	jwtToken, err := jwt.ParseWithClaims(token, &JwtClaim{},
		func(token *jwt.Token) (i interface{}, e error) {
			secret := Config.GetString("jwt.secret")
			return []byte(secret), nil
		})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*JwtClaim); ok && jwtToken.Valid {
			return claim, true
		}
	}
	return nil, false
}

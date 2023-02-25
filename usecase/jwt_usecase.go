package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtUseCase interface {
	GenerateToken(usersId string)(string,error)
	ValidateToken(token string) (*jwt.Token,error)
}

type JwtCustomClaim struct {
	userId string `json:"user_id"`
	jwt.StandardClaims
}

type jwtUseCase struct {
	secretKey string
	issuer string
}
func NewJwtUseCase()JwtUseCase{
	return &jwtUseCase{
		secretKey: getSecretKey(),
		issuer: "kelvin",
	}
}

func getSecretKey() string{
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "kelvinn"
	}
	return secretKey
}

func(j *jwtUseCase)GenerateToken(usersId string)(string, error){
	claims := &JwtCustomClaim{
		 usersId,
		 jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0,0,1).Unix(),
			Issuer: j.issuer,
			IssuedAt: time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t,err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Printf("failed to generate token %v", err)
		panic(err)
	}

	return t,nil
}

func(j *jwtUseCase) ValidateToken(token string) (*jwt.Token,error){
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _,ok := t.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
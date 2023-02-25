package middleware

import (
	"authenctications/usecase"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)
func Authorize(jwtusecase usecase.JwtUseCase) gin.HandlerFunc{
	return func(ctx *gin.Context) {
		autheader := ctx.GetHeader("AUTHORIZATION")
		if autheader == "" {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"err": "failed to authorization"})
			return 
		}
		token,err := jwtusecase.ValidateToken(autheader)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"err": "your token is invalid"})
			return 
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("claims[user_id] :", claims["user_id"])
			log.Println("claims[issuer] :", claims["issuer"])
		}else {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"err": "your token not valid"})
			return 
		}
	}
}
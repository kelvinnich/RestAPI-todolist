package controller

import (
	"authenctications/dto"
	"authenctications/model"
	"authenctications/usecase"
	"authenctications/util"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticationsController interface {
	RegisterController(c *gin.Context)
	LoginController(c *gin.Context)
}

type authencticationsController struct{
	r *gin.Engine
	autusecase usecase.Authentications
	jwt usecase.JwtUseCase
}

func(a *authencticationsController)RegisterController(c *gin.Context){
	var user dto.RegisterUsersDTO
	user.ID = util.NewUUID()
	err := c.ShouldBind(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : err.Error()})
		return
	}
	if !a.autusecase.IsDuplicateEmail(user.Email){
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err" : "failed email is duplicated"})
	}else {
		u := a.autusecase.CreateUsers(user)
		token,_ := a.jwt.GenerateToken(u.ID)
		u.Token = token
		c.JSON(http.StatusOK, u)
	}
}

func (a *authencticationsController) LoginController(c *gin.Context) {
	var loginDTO dto.LoginUsersDTO
	err := c.ShouldBindJSON(&loginDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authResult := a.autusecase.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if authResult == nil {
		log.Println("Invalid credentials")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	user, ok := authResult.(model.Users)
	if !ok {
		log.Println("Error getting user data")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error getting user data"})
		return
	}

	tokenString, err := a.jwt.GenerateToken(user.ID)
	if err != nil {
		log.Println("Error generating token")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	user.Token = tokenString
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "success login", "data": user})
}


func NewAuthenticationController(r *gin.Engine, authusecase usecase.Authentications, jwt usecase.JwtUseCase) *authencticationsController{
	controller := authencticationsController{
		r: r,
		autusecase: authusecase,
		jwt: jwt,
	}

	r.POST("/register", controller.RegisterController)
	r.POST("/login", controller.LoginController)

	return &controller
}


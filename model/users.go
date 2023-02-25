package model


type Users struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"-"`
	Address string `json:"address"`
	Token string `json:"token,omitempty"`
}
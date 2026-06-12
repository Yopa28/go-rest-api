package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role	 string `json:"role"` // admin atau user 
}


type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
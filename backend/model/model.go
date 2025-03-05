package model

type SignupRequest struct {
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	NickName          string `json:"nickName"`
	Gender            string `json:"gender"`
	Age               int    `json:"age"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmedPassword"`
}

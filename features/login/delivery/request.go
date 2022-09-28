package delivery

import "capstone/happyApp/features/login"

type Request struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func toCore(data Request) login.Core {
	return login.Core{
		Email:    data.Email,
		Password: data.Password,
	}
}

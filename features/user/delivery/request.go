package delivery

import "capstone/happyApp/features/user"

type Request struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Gender   string `json:"gender" form:"gender"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Photo    string `json:"photo" form:"photo"`
}

func (req *Request) reqToCore() user.CoreUser {
	return user.CoreUser{
		Name:     req.Name,
		Username: req.Username,
		Gender:   req.Gender,
		Email:    req.Email,
		Password: req.Password,
		Photo:    req.Photo,
	}
}

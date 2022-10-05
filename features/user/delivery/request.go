package delivery

import "capstone/happyApp/features/user"

type Request struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Gender   string `json:"gender" form:"gender"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (req *Request) reqToCore(poto string) user.CoreUser {
	return user.CoreUser{
		Name:     req.Name,
		Username: req.Username,
		Gender:   req.Gender,
		Email:    req.Email,
		Password: req.Password,
		Photo:    poto,
	}
}

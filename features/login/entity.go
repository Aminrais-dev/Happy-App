package login

type Core struct {
	ID       uint
	Email    string
	Password string
	Status   string
}

type UsecaseInterface interface {
	LoginAuthorized(email, password string) (string, error)
}

type DataInterface interface {
	LoginUser(email string) (Core, error)
}

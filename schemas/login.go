package schemas

type Login struct {
	ID       uint
	Email    string `json:"email"`
	Password string `json:"password"`
}

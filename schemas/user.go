package schemas

type User struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	DateCreate      string `json:"datecreate"`
	DateAtualizaton string `json:"dateatualizaton"`
}

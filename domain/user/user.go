package user

type Users struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsVaerified bool   `json:"is_varified"`
}
type UserRepository interface {
	CreateUser(u Users) error
	Login(email string, password string) error
	DeleteUser(uId int) error
}

package user

type Users struct {
	UserId      int    `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsVerified  bool   `json:"is_varified"`
	IsShopOwner bool   `json:"is_shop_owner"`
}
type UserRepository interface {
	CreateUser(u Users) error
	Login(email string, password string) (string, error)
	DeleteUser(uId int) error
}

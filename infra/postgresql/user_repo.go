package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"github/ecommerce/domain/user"
	"github/ecommerce/drivers"
)

// AppError wraps application errors with a code, message, and underlying error
type AppError struct {
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error { return e.Err }

type UserRepo struct {
	DB *sql.DB
}

// compile-time check
var _ user.UserRepository = &UserRepo{}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}
func (r *UserRepo) CreateUser(u user.Users) error {

	// Check if email already exists
	query := `SELECT email FROM users WHERE email=$1;`
	var existingEmail string
	err := r.DB.QueryRow(query, u.Email).Scan(&existingEmail)

	if err != nil && err != sql.ErrNoRows {
		return &AppError{
			Code:    "DB_QUERY_FAIL",
			Message: "failed to check if email exists",
			Err:     err,
		}
	}

	if err == nil {
		return &AppError{
			Code:    "EMAIL_EXISTS",
			Message: "email already registered",
			Err:     nil,
		}
	}

	// Hash password
	pass := drivers.HashSHA256(u.Password)

	// Insert user
	query = `
		INSERT INTO users (name, email, password, is_verified,is_shop_owner)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id;
	`
	var id int
	err = r.DB.QueryRow(query, u.Name, u.Email, pass, u.IsVerified, u.IsShopOwner).Scan(&id)
	if err != nil {
		return &AppError{
			Code:    "DB_INSERT_FAIL",
			Message: "failed to insert user",
			Err:     err,
		}
	}
	query = `INSERT INTO cart(user_id)
	         VALUES ($1)
			 RETURNING cart_id;
	`
	var cart_id int
	err = r.DB.QueryRow(query, id).Scan(&cart_id)
	if err != nil {
		return &AppError{
			Code:    "DB_INSERT_FAIL",
			Message: "failed to insert cart",
			Err:     err,
		}
	}

	fmt.Println("Inserted User ID and Cart ID:", id, cart_id)
	return nil
}

func (r *UserRepo) Login(email, password string) (string, error) {
	query := `SELECT user_id,name,password,is_shop_owner FROM users WHERE email=$1;`
	var userID int
	var Name string
	var storedHash string
	var IsShopOwner bool
	err := r.DB.QueryRow(query, email).Scan(&userID, &Name, &storedHash, &IsShopOwner)

	if errors.Is(err, sql.ErrNoRows) {
		return "", &AppError{
			Code:    "USER_NOT_FOUND",
			Message: "please register first",
			Err:     err,
		}
	} else if err != nil {
		return "", &AppError{
			Code:    "DB_QUERY_FAIL",
			Message: "failed to fetch user by email",
			Err:     err,
		}
	}

	// Hash given password
	pass := drivers.HashSHA256(password)
	if pass != storedHash {
		return "", &AppError{
			Code:    "INVALID_PASSWORD",
			Message: "password not matched",
			Err:     nil,
		}
	}

	// update last login
	updateQuery := `UPDATE users SET last_login = NOW() WHERE user_id = $1;`
	if _, err := r.DB.Exec(updateQuery, userID); err != nil {
		return "", &AppError{
			Code:    "DB_UPDATE_FAIL",
			Message: "failed to update last_login timestamp",
			Err:     err,
		}
	}

	jwt, err := drivers.CreateJWT(userID, Name, IsShopOwner)

	if err != nil || jwt == "" {
		return "", &AppError{
			Code:    "JWT_CREATION",
			Message: "From Login jwt creation not working",
			Err:     nil,
		}
	}

	return jwt, nil
}

func (r *UserRepo) DeleteUser(uId int) error {
	query := `DELETE FROM users WHERE user_id=$1;`
	result, err := r.DB.Exec(query, uId)
	if err != nil {
		return &AppError{
			Code:    "DB_DELETE_FAIL",
			Message: "failed to delete user",
			Err:     err,
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &AppError{
			Code:    "DB_RESULT_FAIL",
			Message: "failed to check delete result",
			Err:     err,
		}
	}
	if rowsAffected == 0 {
		return &AppError{
			Code:    "USER_NOT_FOUND",
			Message: "user not found",
			Err:     nil,
		}
	}
	return nil
}

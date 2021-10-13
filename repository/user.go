package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/egaevan/online-learning/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &User{
		DB: db,
	}
}

func (u *User) FindOne(ctx context.Context, email, password string) (model.User, error) {
	query := `
			SELECT 
				id,
				name,
				email,
				password,
				phone,
			    role
			FROM 
				user
			WHERE
				email = ?`

	user := model.User{}
	err := u.DB.QueryRowContext(ctx, query, email).Scan(
		&user.Id, &user.Name, &user.Email,
		&user.Password, &user.Phone, &user.Role,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("data not found %s", err.Error())
		}
		return user, err
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return user, errors.New("invalid password")
	}

	tk := &model.Token{
		UserID: user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		return user, err
	}

	user.Token = tokenString

	return user, err
}

func (u *User) Fetch(context.Context) error {
	return nil
}

func (u *User) Store(ctx context.Context, user model.User) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(pass)

	query := `
				INSERT INTO user 
					(id, name, email, password, phone, role)
				VALUES
					(?, ?, ?, ?, ?, ?)
			`

	_, err = u.DB.ExecContext(ctx, query,
		user.Id, user.Name, user.Email, user.Password, user.Phone, user.Role)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Update(context.Context) error {
	return nil
}

func (u *User) Delete(ctx context.Context, userID int) error {
	query := `
				UPDATE 
					user
				SET
					flag_aktif = 0
				WHERE
					id = ?
			`

	_, err := u.DB.ExecContext(ctx, query, userID)

	if err != nil {
		return err
	}

	return nil
}

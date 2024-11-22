package user

import (
	"database/sql"
	"fmt"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db}
}

func (service *UserService) Create(user *User) (*User, error) {
	var createdUser User

	fmt.Println(user.UserName, user.Email, user.PasswordHash)
	if err := service.db.QueryRow(`INSERT INTO users (username, email, password_hash) 
									VALUES ($1, $2, $3)
									RETURNING *;`, user.UserName, user.Email, user.PasswordHash).Scan(&createdUser.Id,
		&createdUser.UserName,
		&createdUser.Email,
		&createdUser.PasswordHash); err != nil {
		return nil, fmt.Errorf("failed to create %v: %v", user, err)
	}
	return &createdUser, nil
	// INSERT -> RETURN TO ROW -> COPY IT IN MY TYPE
}

func (service *UserService) Update(user *User) (*User, error) {
	var updatedUser *User
	if err := service.db.QueryRow(`UPDATE users
									SET username = ?,
									email = ?,
									password_hash = ?,
									WHERE id = ? 
									RETURNING id, username, email, password_hash;`, user.UserName, user.Email, user.PasswordHash, user.Id).Scan(updatedUser); err != nil {
		return nil, fmt.Errorf("failed to update %v: %v", user, err)
	}
	return updatedUser, nil
}

func (service *UserService) Delete(userid int) (int, error) {
	if err := service.db.QueryRow(`DELETE FROM users
									WHERE id = ?`, userid); err != nil {
		return -1, fmt.Errorf("failed to create %d: %v", userid, err)
	}
	return userid, nil
}

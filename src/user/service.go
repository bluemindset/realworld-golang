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

func (service *UserService) Create(user User) (User, error) {
	var createdUser User

	fmt.Println(user.UserName, user.Email, user.PasswordHash)
	if err := service.db.QueryRow(`INSERT INTO users (username, email, password_hash) 
									VALUES ($1, $2, $3)
									RETURNING *;`, user.UserName, user.Email, user.PasswordHash).Scan(&createdUser.Id,
		&createdUser.UserName,
		&createdUser.Email,
		&createdUser.PasswordHash); err != nil {
		return User{}, fmt.Errorf("failed to create %v: %v", user, err)
	}
	return createdUser, nil
	// INSERT -> RETURN TO ROW -> COPY IT IN MY TYPE
}

func (service *UserService) Update(user User) (User, error) {
	var updatedUser User
	if err := service.db.QueryRow(`UPDATE users
									SET username = $1,
									email = $2,
									password_hash = $3,
									WHERE id = $4 
									RETURNING id, username, email, password_hash;`, user.UserName, user.Email, user.PasswordHash, user.Id).Scan(&updatedUser); err != nil {
		return User{}, fmt.Errorf("failed to update %v: %v", user, err)
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

func (service *UserService) Read() ([]User, error) {
	var allUsers []User
	rows, err := service.db.Query(`SELECT username, email FROM USERS`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	fmt.Println(rows)
	for rows.Next() {
		var usr User

		if err := rows.Scan(&usr.UserName, &usr.Email); err != nil {
			return []User{}, err
		}
		fmt.Println(usr)
		allUsers = append(allUsers, usr)
	}
	if err = rows.Err(); err != nil {
		return []User{}, err
	}
	return allUsers, nil
}

func (service *UserService) ReadById(userId int) (User, error) {
	var user User
	fmt.Println(userId)
	if err := service.db.QueryRow(`SELECT * FROM USERS WHERE id=$1;`, userId).Scan(&user.Id, &user.UserName, &user.Email, &user.PasswordHash); err != nil {
		return User{}, err
	}

	return user, nil
}

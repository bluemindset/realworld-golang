package user

import (
	"realworld/app/test"
	"testing"
)

// Test driven development
// We first write tests and then code
// Scenarios => function, res

// When writing a db connection function
// 1) Driver by native go for mocking
// 2) Memory Aid

func TestCreateSuccess(t *testing.T) {
	//create new db connection
	//new user service
	//call create
	db := test.DbConnection(t)
	var user = &User{
		Email:        "nick@lambdasec.com",
		UserName:     "nick",
		PasswordHash: "1234",
	}

	userService := NewUserService(db)
	newUser, err := userService.Create(user)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newUser.Email != user.Email {
		t.Errorf("Expected to be equal: %v != %v", newUser.Email, user.Email)
	}
	if newUser.PasswordHash != user.PasswordHash {
		t.Errorf("Expected to be equal: %v != %v", newUser.PasswordHash, user.PasswordHash)
	}
	if newUser.UserName != user.UserName {
		t.Errorf("Expected to be equal: %v != %v", newUser.UserName, user.UserName)
	}
}

func TestCreateError(t *testing.T) {

	db := test.DbConnection(t)

	var user = &User{
		Email:        "nick@lambdasec.com",
		UserName:     "nick",
		PasswordHash: "1234",
	}
	var userDummy = &User{
		Email:        "nick@lambdasec.com1",
		UserName:     "nick1",
		PasswordHash: "12341",
	}

	articleService := NewUserService(db)
	newUser, err := articleService.Create(user)

	if err != nil {
		t.Errorf("Unexpected error happened: %v", err)
	}
	if newUser.Email == userDummy.Email {
		t.Errorf("Expected to be equal: %v != %v", newUser.Email, user.Email)
	}
	if newUser.PasswordHash == userDummy.PasswordHash {
		t.Errorf("Expected to be equal: %v != %v", newUser.PasswordHash, user.PasswordHash)
	}
	if newUser.UserName == userDummy.UserName {
		t.Errorf("Expected to be equal: %v != %v", newUser.UserName, user.UserName)
	}

}

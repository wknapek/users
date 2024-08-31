package usermanager

import (
	"fmt"
	"main/model"
	"sync"
)

var (
	instance *UserManager
	once     sync.Once
)

type UserManager struct {
	users map[string]model.User
}

func GetUserManager() *UserManager {
	once.Do(func() {
		instance = &UserManager{users: make(map[string]model.User)}
	})
	return instance
}

func (u *UserManager) AddUser(user model.User) error {
	_, ok := u.users[user.ID]
	if ok {
		return fmt.Errorf("user %s already exists", user.ID)
	}
	u.users[user.ID] = user
	return nil
}

func (u *UserManager) GetUserByID(userID string) (error, *model.User) {
	usr, ok := u.users[userID]
	if !ok {
		return fmt.Errorf("user with id %s not exist", userID), nil
	}
	return nil, &usr
}

func (u *UserManager) GetAllUsers() []model.User {
	var users []model.User
	for _, user := range u.users {
		users = append(users, user)
	}
	return users
}

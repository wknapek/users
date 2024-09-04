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
	users *sync.Map
}

func GetUserManager() *UserManager {
	once.Do(func() {
		instance = &UserManager{users: new(sync.Map)}
	})
	return instance
}

func (u *UserManager) AddUser(user model.User) error {
	_, ok := u.users.Load(user.ID)
	if ok {
		return fmt.Errorf("user %s already exists", user.ID)
	}
	u.users.Store(user.ID, user)
	return nil
}

func (u *UserManager) GetUserByID(userID string) (error, *model.User) {
	usr, ok := u.users.Load(userID)
	if !ok {
		return fmt.Errorf("user with id %s not exist", userID), nil
	}
	retUsr, ok := usr.(model.User)
	if !ok {
		return fmt.Errorf("error casting user"), nil
	}
	return nil, &retUsr
}

func (u *UserManager) GetAllUsers() []model.User {
	var users []model.User
	u.users.Range(func(k, v interface{}) bool {
		users = append(users, v.(model.User))
		return true
	})
	return users
}

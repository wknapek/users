package usermanager

import (
	"github.com/stretchr/testify/assert"
	"main/model"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	usrMgr := UserManager{new(sync.Map)}
	err := usrMgr.AddUser(model.User{
		ID:         "testID",
		Name:       "TestName",
		SignUpTime: 6,
	})
	assert.NoError(t, err)
	err = usrMgr.AddUser(model.User{
		ID:         "testID",
		Name:       "TestName",
		SignUpTime: 6,
	})
	assert.Error(t, err)
}

func TestGetID(t *testing.T) {
	usrMgr := UserManager{new(sync.Map)}
	err := usrMgr.AddUser(model.User{
		ID:         "testID",
		Name:       "TestName",
		SignUpTime: 6,
	})
	assert.NoError(t, err)
	err, usr := usrMgr.GetUserByID("testID")
	assert.NoError(t, err)
	assert.NotNil(t, usr)
	assert.Equal(t, "testID", usr.ID)
}

func TestGetAll(t *testing.T) {
	usrMgr := UserManager{new(sync.Map)}
	err := usrMgr.AddUser(model.User{
		ID:         "testID",
		Name:       "TestName",
		SignUpTime: 6,
	})
	assert.NoError(t, err)
	err = usrMgr.AddUser(model.User{
		ID:         "testID1",
		Name:       "TestName2",
		SignUpTime: 3,
	})
	assert.NoError(t, err)
	var count int
	usrMgr.users.Range(func(key, value any) bool {
		count++
		return true
	})
	assert.Equal(t, count, 2)
}

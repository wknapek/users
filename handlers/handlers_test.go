package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"main/model"
	"main/security"
	"main/usermanager"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	validLogin := []byte(`{"username": "admin", "password": "test"}`)
	req := httptest.NewRequest("PUT", "/login", bytes.NewBuffer(validLogin))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	security.Login(w, req)
	assert.Equal(t, 200, w.Code)
	testToken, err := io.ReadAll(w.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, testToken)
	validUser := []byte(`{"id": "testID", "name": "test", "signUpTime": 1}`)
	req = httptest.NewRequest("POST", "/create", bytes.NewBuffer(validUser))
	req.Header.Set("Authorization", "Bearer "+string(testToken))
	CreateUserHandler(w, req)
	assert.Equal(t, 200, w.Code)
	usrMgrTest := usermanager.GetUserManager()
	usrMgrTest.AddUser(model.User{
		ID:         "testID",
		Name:       "test",
		SignUpTime: 1,
	})
	req = httptest.NewRequest("GET", "/get?id=testID", nil)
	req.Header.Set("Authorization", "Bearer "+string(testToken))
	GetUserHandler(w, req)
	assert.Equal(t, 200, w.Code)
	bodyBytes, _ := io.ReadAll(w.Body)
	assert.Contains(t, string(bodyBytes), "testID")
	usrMgrTest.AddUser(model.User{
		ID:         "testID2",
		Name:       "test2",
		SignUpTime: 3,
	})
	req = httptest.NewRequest("GET", "/getall", nil)
	req.Header.Set("Authorization", "Bearer "+string(testToken))
	GetAllUsersHandler(w, req)
	assert.Equal(t, 200, w.Code)
	bodyBytesGetAll, _ := io.ReadAll(w.Body)
	users := make([]model.User, 0)
	json.Unmarshal(bodyBytesGetAll, &users)
	assert.Len(t, users, 2)
}

func TestHandlersBadToken(t *testing.T) {
	validLogin := []byte(`{"username": "admin", "password": "test"}`)
	req := httptest.NewRequest("PUT", "/login", bytes.NewBuffer(validLogin))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	security.Login(w, req)
	assert.Equal(t, 200, w.Code)
	testToken, err := io.ReadAll(w.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, testToken)
	validUser := []byte(`{"id": "testID", "name": "test", "signUpTime": 1}`)
	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", "/create", bytes.NewBuffer(validUser))
	req.Header.Set("Authorization", "Bearer "+"badToken")
	CreateUserHandler(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestHandlersBadUser(t *testing.T) {
	validLogin := []byte(`{"username": "test", "password": "test"}`)
	req := httptest.NewRequest("PUT", "/login", bytes.NewBuffer(validLogin))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	security.Login(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	response, err := io.ReadAll(w.Body)
	assert.NoError(t, err)
	assert.Equal(t, string(response), "Invalid credentials")
}

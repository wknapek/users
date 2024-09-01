package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"main/model"
	"main/security"
	"main/usermanager"
	"net/http"
)

var ValidateHandler *validator.Validate

func init() {
	ValidateHandler = validator.New()
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := security.VerifyToken(tokenString)
	if res != http.StatusOK {
		if res == http.StatusUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		if res == http.StatusBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
	}
	u := model.User{}
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Msg("Could not decode body")
		return
	}
	err = ValidateHandler.Struct(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Error().Err(err).Msg("Error validating user")
		return
	}
	mgr := *usermanager.GetUserManager()
	err = mgr.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "user already exists")
		log.Error().Err(err).Msg("Could not add user")
	}
	log.Info().Msg("User created")
	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := security.VerifyToken(tokenString)
	if res != http.StatusOK {
		if res == http.StatusUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		if res == http.StatusBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
	}
	id := r.URL.Query().Get("id")
	mgr := *usermanager.GetUserManager()
	err, usr := mgr.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "user not found")
		return
	}
	bytes, err := json.Marshal(usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Could not encode body")
		return
	}
	fmt.Fprint(w, string(bytes))
	w.WriteHeader(http.StatusOK)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res := security.VerifyToken(tokenString)
	if res != http.StatusOK {
		if res == http.StatusUnauthorized {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}
		if res == http.StatusBadRequest {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
	}
	mgr := *usermanager.GetUserManager()
	users := mgr.GetAllUsers()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not encode users")
		log.Error().Err(err).Msg("could not encode users")
	}
	w.WriteHeader(http.StatusOK)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"main/model"
	"main/security"
	"main/usermanager"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := security.VerifyToken(w, r)
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
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Info().Msg("Could not decode body")
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
	res := security.VerifyToken(w, r)
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
	id := chi.URLParam(r, "id")
	mgr := *usermanager.GetUserManager()
	err, usr := mgr.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "user not found")
		return
	}
	err = json.NewEncoder(w).Encode(usr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not encncode user")
		log.Error().Err(err).Msg("could not encode user")
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := security.VerifyToken(w, r)
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

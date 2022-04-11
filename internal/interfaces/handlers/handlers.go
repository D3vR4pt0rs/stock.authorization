package handlers

import (
	"encoding/json"
	"net/http"

	"authentication/internal/entities"
	"authentication/internal/usecases/storage"

	"github.com/D3vR4pt0rs/logger"

	"github.com/gorilla/mux"
)

func registration(app storage.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Got new request for registration profile")
		errorMessage := "Error creating profile"

		newCredentials := entities.Credentials{}

		err := json.NewDecoder(r.Body).Decode(&newCredentials)
		if err != nil {
			logger.Error.Printf("Failed to decode credentials. Got %v", err)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		err = app.Registration(newCredentials)
		switch err {
		case storage.AccountExistError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case storage.InternalError:
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]interface{})
		resp["status"] = "success"
		json.NewEncoder(w).Encode(resp)
	})
}

func signup(app storage.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error creating profile"
		credentials := entities.Credentials{}

		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			logger.Error.Printf("Failed to decode credentials. Got %v", err)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		logger.Info.Println("Got new request for authorization to account: ", credentials.Email)

		token, err := app.SignUp(credentials)

		switch err {
		case storage.AccountNotFoundError, storage.WrongPasswordError:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case storage.InternalError:
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]interface{})
		resp["token"] = token

		json.NewEncoder(w).Encode(resp)
	})
}

func Make(r *mux.Router, app storage.Controller) {
	apiUri := "/api"
	serviceRouter := r.PathPrefix(apiUri).Subrouter()
	serviceRouter.Handle("/account/registration", registration(app)).Methods("POST")
	serviceRouter.Handle("/account/signup", signup(app)).Methods("POST")
}

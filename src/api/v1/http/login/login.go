package login

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/tchtsk/treatfield-api/src/api/v1/http/responses"
	"github.com/tchtsk/treatfield-api/src/mysql/users"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var loginRequest loginRequest
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&loginRequest)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(loginRequest.Email)

	user, err := users.GetUserByEmailAndPassword(loginRequest.Email, loginRequest.Password)

	if err != nil {
		log.Println(err)
		http.Error(response, err.Error(), http.StatusNotFound)
		return
	}

	responseData := &responses.UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: time.Unix(user.CreatedAt, 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(user.UpdatedAt, 0).Format(time.RFC3339),
	}

	json.NewEncoder(response).Encode(responseData)
}

func FirstUserHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := users.GetUserByEmailAndPassword("vtchtsk@gmail.com", "vtchtsk@gmail.com")

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := &responses.UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: time.Unix(user.CreatedAt, 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(user.UpdatedAt, 0).Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(response)
}

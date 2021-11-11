package create

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/tchtsk/treatfield-api/src/api/v1/http/responses"
	"github.com/tchtsk/treatfield-api/src/mysql/users"
)

type createRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUserHandler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var requestPayload createRequest
	decoder := json.NewDecoder(request.Body)

	err := decoder.Decode(&requestPayload)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := users.CreateUserByEmailAndPassword(requestPayload.Email, requestPayload.Password)

	if err != nil {
		log.Println(err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := &responses.UserResponse{
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: time.Unix(user.CreatedAt, 0).Format(time.RFC3339),
		UpdatedAt: time.Unix(user.UpdatedAt, 0).Format(time.RFC3339),
	}

	json.NewEncoder(response).Encode(responseData)
}

func ClearUserHandler(response http.ResponseWriter, request *http.Request) {
	users.Clear()
}

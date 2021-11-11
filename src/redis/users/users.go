package users

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/tchtsk/treatfield-api/src/models"
	"github.com/tchtsk/treatfield-api/src/redis"
)

var cacheAccessMutex sync.Mutex

func AddUserToSave(email string, password string) (models.User, int) {
	cacheAccessMutex.Lock()

	defer cacheAccessMutex.Unlock()

	model := models.User{
		Email:     email,
		Username:  email,
		Password:  password,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	currentValue, err := redis.GetInstance().Get("users").Result()
	if err != nil {
		log.Println("No data in cache", err)
	}

	var currentValuesMap []models.User
	if err != nil || currentValue != "" {
		json.Unmarshal([]byte(currentValue), &currentValuesMap)
		log.Println("Models count in cache", len(currentValuesMap))
	}
	err = nil

	currentValuesMap = append(currentValuesMap, model)
	updatedValuesMapJson, err := json.Marshal(currentValuesMap)

	if err != nil {
		log.Println(err)
		return models.User{}, 0
	}

	err = redis.GetInstance().Set("users", updatedValuesMapJson, 0).Err()

	if err != nil {
		log.Println(err)
		return models.User{}, 0
	}

	len := len(currentValuesMap)

	return model, len
}

func GetAndFlush() []models.User {
	cacheAccessMutex.Lock()
	defer cacheAccessMutex.Unlock()

	currentValue, err := redis.GetInstance().Get("users").Result()
	if err != nil {
		log.Println("No data in cache", err)
		return []models.User{}
	}

	var currentValuesMap []models.User
	json.Unmarshal([]byte(currentValue), &currentValuesMap)

	Flush()

	return currentValuesMap
}

func Flush() {
	redis.GetInstance().Set("users", "", 0)
}

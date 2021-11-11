package users

import (
	"database/sql"
	"log"
	"math"
	"math/rand"
	"sync"

	"github.com/tchtsk/treatfield-api/src/models"
	"github.com/tchtsk/treatfield-api/src/mysql"
	"github.com/tchtsk/treatfield-api/src/redis/users"
)

var createUserMutex sync.Mutex

func GetUserByEmailAndPassword(email string, password string) (models.User, error) {
	var user models.User

	stmt, err := mysql.MysqlDb.Prepare("SELECT email,username,UNIX_TIMESTAMP(created),UNIX_TIMESTAMP(updated) FROM tmp WHERE email=? AND password=? LIMIT 1")
	if err != nil {
		log.Println(err)
	}

	defer stmt.Close()

	var raw *sql.Row = stmt.QueryRow(email, password)

	if raw.Err() != nil {
		log.Println(raw.Err())
	}

	err = raw.Scan(
		&user.Email,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

func CreateUserByEmailAndPassword(email string, password string) (models.User, error) {
	userModel, modelsCountInCache := users.AddUserToSave(email, password)

	if float64(modelsCountInCache)-float64(rand.Intn(20)+1)*math.Log(rand.Float64()) >= 500.0 {
		stmt, err := mysql.MysqlDb.Prepare("INSERT INTO tmp (email,username,password,created,updated) VALUES (?,?,?,?,?)")
		defer stmt.Close()

		if err != nil {
			log.Println(err, "Prepare statement error")
			return models.User{}, err
		}

		for _, v := range users.GetAndFlush() {
			_, err = stmt.Exec(
				v.Email,
				v.Email,
				v.Password,
				v.CreatedAt,
				v.UpdatedAt,
			)

			if err != nil {
				log.Println(err, "Exec statement error")
			} else {
				log.Println("Raw inserted", v)
			}
		}
	}

	return userModel, nil
}

func Clear() {
	users.Flush()
}

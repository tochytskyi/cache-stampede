package users

import (
	"database/sql"
	"log"

	"github.com/tchtsk/treatfield-api/src/mysql"
)

type User struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	CreatedAt int64  `json:"created"`
	UpdatedAt int64  `json:"updated"`
}

func GetUserByEmailAndPassword(email string, password string) (User, error) {
	var user User

	var raw *sql.Row = mysql.
		MysqlDb.
		QueryRow("SELECT id,email,username,created,updated FROM user WHERE email=\"" + email + "\" LIMIT 1")

	if raw.Err() != nil {
		log.Println(raw.Err())
	}

	err := raw.Scan(
		&user.Id,
		&user.Email,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return user, err
}

package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tchtsk/treatfield-api/src/api/v1/http/login"
	create "github.com/tchtsk/treatfield-api/src/api/v1/http/users"
	"github.com/tchtsk/treatfield-api/src/mysql"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "5000"
	}

	mysql.MysqlDb = mysql.Init()
	defer mysql.MysqlDb.Close()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/api/v1/login", login.LoginHandler)
	http.HandleFunc("/api/v1/users", create.CreateUserHandler)
	http.HandleFunc("/api/v1/users/clear", create.ClearUserHandler)

	http.ListenAndServe(":"+httpPort, nil)
}

package main

import (
	"casbin/delivery"
	"casbin/middlewares"
	"casbin/repository"
	"casbin/usecase"
	"database/sql"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func connect() *sql.DB {
	var dbHost, dbUser, dbPass, dbName string
	dbHost = viper.GetString("database.host")
	dbUser = viper.GetString("database.user")
	dbPass = viper.GetString("database.pass")
	dbName = viper.GetString("database.db")
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := connect()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Duration(viper.GetInt("database.maxLifeTime")) * time.Minute)
	db.SetMaxIdleConns(viper.GetInt("database.maxIdleCons"))
	db.SetMaxOpenConns(viper.GetInt("database.maxOpenCons"))

	e := echo.New()

	r, err := casbin.NewEnforcer("auth_model.conf", "policy.csv")
	if err != nil {
		panic(err.Error())
	}
	enforcer := middlewares.Enforcer{Enforcer: r}
	timeOut := time.Duration(viper.GetInt("context.timeout")) * time.Second
	repo := repository.NewRepositoryConfig(db)
	us := usecase.NewUsecaseConfig(repo, timeOut)
	delivery.NewHandler(e, us)
	// e.Use(enforcer.Enforce)
	e.Use(echo.WrapMiddleware(enforcer.Enforce))
	e.Use(echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))

	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}

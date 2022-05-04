package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"log"
)

type Config struct {
	IP   string
	Port int
}

type DB struct {
	config *Config
}

type UserRepository struct {
	db *DB
}

type UserService struct {
	userRepo *UserRepository
}

type UserController struct {
	userService *UserService
}

func (s *UserController) GetUser(ctx *gin.Context) {
	return
}

func NewConfig() (*Config, error) {
	fmt.Print("new Config\n")
	return new(Config), nil
}

func NewDatabase(c *Config) (*DB, error) {
	log.Printf("new DB, need config:%+v\n", c)
	return new(DB), nil
}

func NewUserRepository(db *DB) (*UserRepository, func(), error) {

	log.Printf("new UserRepository, need DB:%+v\n", db)
	x := &UserRepository{
		db: db,
	}
	return x, x.Stop, nil
}

func (s *UserRepository) Stop() {

	log.Println("stop some task ,eg: mq consumer")
	return
}

func NewUserService(UserRepo *UserRepository) (*UserService, error) {
	log.Printf("new NewUserService, need UserRepository:%+v\n", UserRepo)
	return &UserService{
		userRepo: UserRepo,
	}, nil

}

func NewUserController(userService *UserService) (*UserController, error) {
	log.Printf("new NewController, need UserService:%+v\n", userService)
	return &UserController{
		userService: userService,
	}, nil
}

type HttpServer struct {
	userCtrl *UserController
}

func NewHttpServer(userCtrl *UserController) *HttpServer {

	return &HttpServer{
		userCtrl: userCtrl,
	}
}

func (s *HttpServer) Run() error {

	engine := gin.Default()
	engine.GET("/user", s.userCtrl.GetUser)

	return engine.Run()
}

func InitApp() (*HttpServer, func(), error) {
	panic(wire.Build(NewConfig, NewDatabase, NewUserRepository, NewUserService, NewUserController, NewHttpServer))
}

func main() {
	app, clean, err := InitApp()

	if err != nil {
		panic(err)
	}
	defer clean()
	app.Run()

}

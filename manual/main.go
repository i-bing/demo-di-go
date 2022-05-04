package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Config struct {
	IP   string
	Port int
}

type DB struct {
	conf *Config
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

func NewConfig() *Config {
	log.Print("new Config\n")
	return new(Config)
}

func NewDatabase(c *Config) *DB {
	log.Printf("new DB, need conf:%+v\n", c)
	return new(DB)
}

func NewUserRepository(db *DB) *UserRepository {
	log.Printf("new UserRepository, need DB:%+v\n", db)
	return &UserRepository{
		db: db,
	}
}

func NewUserService(UserRepo *UserRepository) *UserService {
	log.Printf("new NewUserService, need UserRepository:%+v\n", UserRepo)
	return &UserService{
		userRepo: UserRepo,
	}

}

func NewUserController(userService *UserService) *UserController {
	log.Printf("new NewController, need UserService:%+v\n", userService)
	return &UserController{
		userService: userService,
	}
}

func NewHttpServer(userCtrl *UserController) *gin.Engine {

	engine := gin.Default()
	engine.GET("/user", userCtrl.GetUser)

	return engine
}

func main() {
	config := NewConfig()
	db := NewDatabase(config)
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userCtrl := NewUserController(userService)
	server := NewHttpServer(userCtrl)
	server.Run()

}

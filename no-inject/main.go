package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Config struct {
	IP   string
	Port int
}

type DataBase struct {
}

type UserRepository struct {
}

func (s *UserRepository) GetUser(id int64) {
	log.Println("user DB", DB)
	return
}

type UserService struct {
}

func (s *UserService) GetUser(id int64) {
	UserRepo.GetUser(id)
	return
}

type UserController struct {
}

func (s *UserController) GetUser(ctx *gin.Context) {
	UserServ.GetUser(int64(100))
	return
}

func NewConfig() *Config {
	log.Print("new Config\n")
	return new(Config)
}

func NewDatabase() *DataBase {
	serverUri := fmt.Sprintf("%s:%d", Conf.IP, Conf.Port)
	log.Printf("new DataBase, need config:%+v\n", serverUri)
	return new(DataBase)
}

func NewUserRepository() *UserRepository {
	return &UserRepository{

	}
}

func NewUserService() *UserService {
	return &UserService{
	}

}

func NewUserController() *UserController {
	return &UserController{
	}
}

func NewHttpServer() *gin.Engine {

	engine := gin.Default()
	engine.GET("/user", UserCtrl.GetUser)

	return engine
}

var Conf *Config
var DB *DataBase
var UserRepo *UserRepository
var UserServ *UserService
var UserCtrl *UserController

func main() {
	Conf = NewConfig()
	DB = NewDatabase()
	UserRepo = NewUserRepository()
	UserServ = NewUserService()
	UserCtrl = NewUserController()
	server := NewHttpServer()
	server.Run()

}

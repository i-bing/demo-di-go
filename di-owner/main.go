package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Config struct {
	IP   string
	Port int
}

type Database struct {
	Conf *Config `di:"config""`
}

type UserRepository struct {
	DB *Database `di:"database"`
}

type UserService struct {
	UserRepo *UserRepository `di:"userRepository"`
}

type UserController struct {
	UserServ *UserService `di:"userService"`
}

func (s *UserController) GetUser(ctx *gin.Context) {
	return
}

type HttpServer struct {
	UserCtrl *UserController `di:"userController"`
}

func NewConfig() *Config {
	log.Print("new Config\n")
	return new(Config)
}

func (s *HttpServer) Run() error {

	engine := gin.Default()
	engine.GET("/user", s.UserCtrl.GetUser)

	log.Println("http inject", s.UserCtrl)
	return engine.Run()
}

func BuildDIContainer() error {
	di := NewContainer()
	di.SetSingleton("config", NewConfig())
	di.SetSingleton("database", &Database{})
	di.SetSingleton("userRepository", &UserRepository{})
	di.SetSingleton("userService", &UserService{})
	di.SetSingleton("userController", &UserController{})
	di.SetSingleton("httpServer", &HttpServer{})

	di.SetPrototype("map", func() (interface{}, error) {
		return make(map[interface{}]interface{}), nil
	})

	if v, err := di.Ensure(di.GetSingleton("database")); err != nil {
		log.Println(err, ", failed to ensure Database", v)
	}
	if v, err := di.Ensure(di.GetSingleton("userRepository")); err != nil {
		log.Println(err, ", failed to ensure UserRepository", v)
	}
	if v, err := di.Ensure(di.GetSingleton("userService")); err != nil {
		log.Println(err, ", failed to ensure UserService", v)
	}
	if v, err := di.Ensure(di.GetSingleton("userController")); err != nil {
		log.Println(err, ", failed to ensure UserController", v)
	}
	if v, err := di.Ensure(di.GetSingleton("httpServer")); err != nil {
		log.Println(err, ", failed to ensure httpServer", v)
	}

	/*	// try to fix import cycle
		container.SetSingleton(definitions.DIServerPush, &impl.ServerPushImpl{})

		if v, err := container.Ensure(&impl.ServerPushImpl{}); err == nil {
			container.SetSingleton(definitions.DIServerPush, v)
		} else {
			log.Errorln(err, ", failed to ensure serverPush")
		}*/

	log.Println(di.String())
	return nil
}
func main() {
	err := BuildDIContainer()

	if err != nil {
		panic(err)
	}
	v := GetContainer().GetSingleton("httpServer")
	v.(*HttpServer).Run()
}

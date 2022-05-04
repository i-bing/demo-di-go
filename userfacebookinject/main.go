package main

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"log"
)

type Config struct {
	IP   string
	Port int
}

type Database struct {
	Conf *Config `inject:""` //每一个需要注入的字段都需要打上 inject:"" 这样的 tag
}

type UserRepository struct {
	DB *Database `inject:""`
}

type UserService struct {
	UserRepo *UserRepository `inject:""`
}

type UserController struct {
	UserService *UserService `inject:""`
}

func (s *UserController) GetUser(ctx *gin.Context) {
	return
}

type HttpServer struct {
	UserCtrl *UserController `inject:""`
}

func NewConfig() *Config {
	log.Print("new Config\n")
	return new(Config)
}

func (s *HttpServer) Run() error {

	log.Println("user ctrl", s.UserCtrl)
	engine := gin.Default()
	engine.GET("/user", s.UserCtrl.GetUser)

	return engine.Run()
}

func main() {
	var s HttpServer

	graph := inject.Graph{} //创建一个 graph 对象, graph 对象将负责管理和注入所有的对象


	err := graph.Provide(
		&inject.Object{
			Value:    NewConfig(),
			Name:     "",
			Complete: false,
			Fields:   nil,
		}, &inject.Object{
			Value: &s,
		})// graph.Provide() 将需要注入的对象提供给 graph

	if err != nil {
		panic(err)
	}
	err = graph.Populate() //调用 Populate 函数，开始进行注入

	if err != nil {
		panic(err)
	}

	s.Run()
}

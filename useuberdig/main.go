package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
)

type Config struct {
	IP   string
	Port int
}

type Database struct {
	config *Config
}


type UserRepository struct {
	db *Database
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

func NewDatabase(c *Config) *Database {
	log.Printf("new Database, need config:%+v\n", c)
	return new(Database)
}

func NewUserRepository(db *Database) *UserRepository {
	log.Printf("new UserRepository, need Database:%+v\n", db)
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

func BuildContainer() *dig.Container {
	container := dig.New()

	container.Provide(NewConfig) //被依赖方的创建方式注册到Provide中
	container.Provide(NewDatabase)
	container.Provide(NewUserRepository)
	container.Provide(NewUserService)
	container.Provide(NewUserController)
	container.Provide(NewHttpServer)

	b := &bytes.Buffer{}
	if err := dig.Visualize(container, b); err != nil {
		panic(err)
	}
	fmt.Println(b.String())
	return container
}

func main() {
	container := BuildContainer()
	fmt.Println(container.String())

	err := container.Invoke(func(server *gin.Engine) {
		server.Run() //服务在启动的时候，只需要在Invoke中启动就可以了
	})
	if err != nil {
		panic(err)
	}
}

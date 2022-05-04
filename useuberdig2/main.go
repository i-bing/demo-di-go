package main
/*
参数对象 https://pkg.go.dev/go.uber.org/dig#hdr-Parameter_Objects
结果对象 https://pkg.go.dev/go.uber.org/dig#hdr-Result_Objects
*/
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

type MsDB struct {
	config *Config
}

type Cache struct {
	config *Config
}

type ES struct {
	config *Config
}
type Mongo struct {
	config *Config
}

type DataSource struct {
	dig.Out
	db    *MsDB
	cache *Cache
	es    *ES
	mongo *Mongo
}
type UserRepository struct {
	db    *MsDB
	cache *Cache
}
type OrderRepository struct {
	db    *MsDB
	cache *Cache

	es    *ES
	mdb *Mongo
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

//result obejcts
func NewDataSource(c *Config) (*MsDB, *Cache, *ES, *Mongo) {
	log.Printf("new Database, need config:%+v\n", c)
	return new(MsDB), new(Cache), new(ES), new(Mongo)
}

func NewUserRepository(db *MsDB, cache *Cache) *UserRepository {
	log.Printf("new UserRepository, need Database:%+v\n", db)
	return &UserRepository{
		db:    db,
		cache: cache,
	}
}

/*
func NewOrderRepository(db *MsDB,cache *Cache,es *ES,mdb *Mongo) *OrderRepository {
	log.Printf("new UserRepository, need Database:%+v\n", param)
	return &OrderRepository{
		db:    db,
		cache: cache,
		es:   es,
		mdb: mdb,
	}
}*/

// param objects
type UserRepoParam struct {
	dig.In
	*MsDB
	*Cache
	*ES
	*Mongo
}

func NewOrderRepository(param *UserRepoParam) *OrderRepository {
	log.Printf("new UserRepository, need Database:%+v\n", param)
	return &OrderRepository{
		db:    param.MsDB,
		cache: param.Cache,
		es:    param.ES,
		mdb: param.Mongo,
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

	container.Provide(NewConfig)
	container.Provide(NewDataSource)
	container.Provide(NewUserRepository)
	container.Provide(NewOrderRepository)
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
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	controllerFactory := NewControllerFactory(r)
	userController := controllerFactory.NewUserController()

	userController.attachRoutes()

	r.Run()
}

type Config struct {
	Database struct {
		connectionString string
	}
	Server struct {
		port string
	}
	logLevel string
}

func GetConfig() *Config {
	config := &Config{}
	connectionString := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	logLevel := os.Getenv("LOG_LEVEL")
	for _, e := range []string{connectionString, port, logLevel} {
		if e == "" {
			panic("Missing environment variable")
		}
	}
	config.Database.connectionString = connectionString
	config.Server.port = port
	config.logLevel = logLevel
	return config
}

func NewControllerFactory(r *gin.Engine) *ControllerFactory {
	config := GetConfig()
	db, err := sql.Open("postgres", config.Database.connectionString)
	if err != nil {
		panic(err)
	}
	return &ControllerFactory{
		r:  r,
		db: db,
	}
}

type ControllerFactory struct {
	r  *gin.Engine
	db *sql.DB
}

func (f *ControllerFactory) NewUserController() *UserController {
	userService := UserService{
		db: f.db,
	}
	return NewUserController(f.r, &userService)
}

func NewUserController(r *gin.Engine, userService *UserService) *UserController {
	controller := &UserController{
		userService: userService,
		router:      r.Group("/user"),
	}

	return controller
}

func (c *UserController) attachRoutes() {
	c.router.GET("/:id", c.getUser)
	c.router.POST("/", c.createUser)
}

func (controller *UserController) getUser(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Hello %s", c.Param("id")))
}

func (controller *UserController) createUser(c *gin.Context) {
}

func (s *UserService) provisionToken() string {
	return ""
}

func (s *UserService) createUser(u string, p string) {
}

type CreateUserDTO struct {
	username string
	password string
	email    string
}

type UserResponseDTO struct {
	username string
	email    string
}

type LoginUserDTO struct {
	username string
	password string
}

type UserService struct {
	db *sql.DB
}

type IUserService interface {
	provisionToken() string
}

type UserController struct {
	userService *UserService
	router      *gin.RouterGroup
}

type IAuthService interface {
	loginUser(string, string)
	createToken(string) string
}

type AuthController struct {
	authService *IAuthService
	router      *gin.RouterGroup
}

type AuthService struct{}

package main

import (
	"fmt"
	"log"

	"github.com/ahmetk3436/pkg/api"
	"github.com/ahmetk3436/pkg/repository/post"
	"github.com/ahmetk3436/pkg/service"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// App is alias for api.App{}
type App struct {
	e  *echo.Echo
	DB *gorm.DB
}

func main() {
	a := App{}

	// Initialize storage
	a.initialize(
		"34.155.128.105",
		"3306",
		"root",
		"r)qDpDcl~$mkcs&Z",
		"pomodoro")

	// Initialize routes
	a.routes()

	// Start server
	a.run(":8010")
}

func (a *App) initialize(host, port, username, password, dbname string) {
	var err error

	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, host, port, dbname)

	a.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// repository := post.NewRepository(a.DB)
	a.e = echo.New()
}

func (a *App) run(addr string) {
	fmt.Printf("Server started at %s\n", addr)
	a.e.Logger.Fatal(a.e.Start(addr))
}

func (a *App) routes() {
	postAPI := InitPostAPI(a.DB)
	a.e.GET("/examDate/:id", postAPI.FindByID)
	a.e.POST("/examDate/", postAPI.CreatePost)
}

// InitPostAPI ..
func InitPostAPI(db *gorm.DB) api.PostAPI {
	postRepository := post.NewRepository(db)
	postService := service.NewPostService(postRepository)
	postAPI := api.NewPostAPI(postService)
	postAPI.Migrate()
	return postAPI
}

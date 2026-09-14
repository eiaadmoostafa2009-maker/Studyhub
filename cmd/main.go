package main

import (
	
	"log"
	"net/http"
	"os"
	"studyhub/internal/config"
	courseHandler "studyhub/internal/handlers/course"
	postHandler "studyhub/internal/handlers/post"
	userHandler "studyhub/internal/handlers/user"
	courseRepository "studyhub/internal/repository/course"
	postRepository "studyhub/internal/repository/post"
	userRepo "studyhub/internal/repository/user"
	courseServe "studyhub/internal/service/course"
	postServe "studyhub/internal/service/post"
	userServe "studyhub/internal/service/user"
	"studyhub/pkg/internalsql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
func main(){
	r :=gin.Default()
	validate :=validator.New()
    
    // Load configuration
	cfg, err := config.LoadConfig()
	if err != nil{
		log.Fatal(err)
	}
    // Connect to the database
	db, err := internalsql.DBConnect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	//initialize frontend
	r.LoadHTMLGlob("internal/template/*.html")
	r.Static("/static", "./internal/static")
	r.Static("/js", "./internal/static/js")
	r.Static("/css", "./internal/static/css")
	r.Static("/uploads", "./uploads")

	//index route
	r.GET("/", func (c *gin.Context){
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Initialize the repositories
	userRepo := userRepo.NewUserRepository(db)
	postRepo := postRepository.NewPostRepository(db)
	courseRepo := courseRepository.NewCourseRepository(db)
	
	
	// Initialize the services
	userService := userServe.NewUserService(cfg, userRepo)
	postService := postServe.NewPostService(cfg, postRepo)
	courseService := courseServe.NewCourseService(cfg, courseRepo)
	

	// Initialize the handlers
	userHandler := userHandler.NewUserHandler(r, validate, userService)
	postHandler := postHandler.NewPostHandler(r, validate, postService)
    courseHandler := courseHandler.NewCourseHandler(r, validate, courseService)
	
	//initialize routesfor handlers
	userHandler.RegisterRoutes(cfg.SecretKey, userRepo)
    postHandler.RouteList(cfg.SecretKey, userRepo)
    courseHandler.RouteList(cfg.SecretKey, userRepo)
    
	port := os.Getenv("PORT")
    if port == "" {
       port = "8080"
    }

    r.Run(":" + port)
}
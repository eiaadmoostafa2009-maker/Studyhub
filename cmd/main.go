package main

import (
	"log"
	"net/http"
	"os"
	"studyhub/internal/config"
	courseHandler "studyhub/internal/handlers/course"
	postHandler "studyhub/internal/handlers/post"
	userHandler "studyhub/internal/handlers/user"
	"studyhub/internal/integrations/telegram"
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

func main() {
	r := gin.Default()
	validate := validator.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := internalsql.DBConnect(cfg)
	if err != nil {
		log.Fatal(err)
	}


	r.LoadHTMLGlob("internal/template/*.html")
	r.Static("/static", "./internal/static")
	r.Static("/js", "./internal/static/js")
	r.Static("/css", "./internal/static/css")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	userRepo := userRepo.NewUserRepository(db)
	postRepo := postRepository.NewPostRepository(db)
	courseRepo := courseRepository.NewCourseRepository(db)
   
	client := telegram.NewClient()

	userService := userServe.NewUserService(cfg, userRepo)
	postService := postServe.NewPostService(cfg, postRepo)
	courseService := courseServe.NewCourseService(cfg, courseRepo, client)

	status, err := userService.SeedAdmin()
	if err != nil {
		log.Fatal("failed to seed admin:", err)
	}
	log.Println("admin seed completed with status:", status)

	userHandler := userHandler.NewUserHandler(r, validate, userService)
	postHandler := postHandler.NewPostHandler(r, validate, postService)
	courseHandler := courseHandler.NewCourseHandler(r, validate, courseService)

	userHandler.RegisterRoutes(cfg.SecretKey, userRepo)
	postHandler.RouteList(cfg.SecretKey, userRepo)
	courseHandler.RouteList(cfg.SecretKey, userRepo)

	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Port
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}

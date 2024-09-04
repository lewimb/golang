package main

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository/authRepository"
	dbRepository "a21hc3NpZ25tZW50/repository/dbRepository"
	"errors"

	// "a21hc3NpZ25tZW50/service"
	"fmt"
	"log"
	"net/http"

	_ "embed"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

func AuthMiddleware(authRepo *authRepository.Repository) gin.HandlerFunc {
	// TODO: answer here
	return nil
}

func Connect(creds *Credential) (*gorm.DB, error) {
	// this is only an example, please modify it to your need
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	// connect using gorm
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func SetupRouter(dbRepo *dbRepository.Repository, authRepo *authRepository.Repository) *gin.Engine {
	// svc := service.NewService(*dbRepo, authRepo)

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusBadRequest, "Bad request") // TODO: replace this
	})

	router.POST("/users", func(c *gin.Context) {
		var signUp model.User
		err := c.ShouldBindJSON(&signUp)
		if err != nil {
			c.String(http.StatusBadRequest, "invalid username or password")
		}

		if _, err = dbRepo.GetUserByUsername(signUp.Username); err == nil {
			c.String(http.StatusConflict, "username already registered")
			return
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			c.String(http.StatusInternalServerError, "could not check user existence")
			return
		}

		newUser, err := dbRepo.AddUser(signUp)
		if err != nil {
			c.String(http.StatusInternalServerError, "could not check user existence")
			return
		}

		c.JSON(http.StatusCreated, newUser)
	})

	router.POST("/signin", func(c *gin.Context) {
		// TODO: answer here
	})

	// Apply the auth middleware to routes that need protection
	protected := router.Group("/")
	protected.Use(AuthMiddleware(authRepo))
	{
		protected.GET("/photos", func(c *gin.Context) {
			// TODO: answer here
		})

		protected.GET("/photos/:id", func(c *gin.Context) {
			// TODO: answer here
		})

		protected.POST("/photos", func(c *gin.Context) {
			// TODO: answer here
		})

		protected.DELETE("/photos/:id", func(c *gin.Context) {
			// TODO: answer here
		})
	}

	// Catch-all route for undefined routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
	})

	return router
}

func main() {
	dbCredential := Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "postgres",
		DatabaseName: "kampusmerdeka",
		Port:         5432,
	}
	dbConn, err := Connect(&dbCredential)
	if err != nil {
		log.Fatal(err)
	}

	if err = dbConn.Migrator().DropTable("users", "photos"); err != nil {
		log.Fatal("failed droping table:" + err.Error())
	}

	if err = dbConn.AutoMigrate(&model.User{}, &model.Photo{}); err != nil {
		log.Fatal("failed migrating table:" + err.Error())
	}

	authRepo := authRepository.NewRepository()
	dbRepo := dbRepository.NewRepository(dbConn)

	router := SetupRouter(dbRepo, authRepo)

	router.Run()

}

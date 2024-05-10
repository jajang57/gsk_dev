package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}
type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (server *Server) Initialize(appConfig AppConfig, dbconfig DBConfig) {
	fmt.Println("wlcome to gsk_dev " + appConfig.AppName)

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbconfig.DBHost, dbconfig.DBUser, dbconfig.DBPassword, dbconfig.DBName, dbconfig.DBPort)
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed connect to DB")
	}

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Printf("listen to port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
func Run() {
	var server = Server{}
	var appConfig = AppConfig{}

	var dbConfig = DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env.example file")
	}

	appConfig.AppName = getEnv("APP_NAME", "gogsk")
	appConfig.AppEnv = getEnv("APP_ENV", "development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	dbConfig.DBName = getEnv("DB_NAME", "dbname")
	dbConfig.DBPort = getEnv("DB_PORT", "5433")

	server.Initialize(appConfig, dbConfig)
	server.Run(":" + appConfig.AppPort)
}

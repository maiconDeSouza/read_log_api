package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/maiconDeSouza/read_log_api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Env struct {
	VersionAPI string
	DBName     string
	DBUser     string
	DBPass     string
	DBPort     string
	ServerPort string
	SecretKey  string
}

func initDotenv() *Env {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌​ Erro ao carregar o arquivo .env:", err)
	}

	env := &Env{
		VersionAPI: os.Getenv("VERSION_API"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPass:     os.Getenv("DB_PASS"),
		DBPort:     os.Getenv("DB_PORT"),
		ServerPort: os.Getenv("SERVER_PORT"),
		SecretKey:  os.Getenv("SECRET_KEY"),
	}
	return env
}

func initMUX() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}

func initDB(env *Env) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		env.DBUser,
		env.DBPass,
		env.DBName,
		env.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Falha ao conectar no banco de dados: %v", err)
	}

	user := models.User{}
	err = db.AutoMigrate(&user)
	if err != nil {
		log.Fatalf("❌ Falha ao migrar movie: %v", err)
	}

	userBook := models.UserBook{}
	err = db.AutoMigrate(&userBook)
	if err != nil {
		log.Fatalf("❌ Falha ao migrar movie: %v", err)
	}

	book := models.Book{}
	err = db.AutoMigrate(&book)
	if err != nil {
		log.Fatalf("❌ Falha ao migrar movie: %v", err)
	}

	fmt.Println("✅ Conectado ao banco de dados com sucesso!")
	return db
}

func InitConfig() (*Env, *http.ServeMux, *gorm.DB) {
	env := initDotenv()
	mux := initMUX()
	db := initDB(env)

	return env, mux, db
}

package main

import (
	"database/sql"
	"fmt"
	"github/Himesh-Kundal/PropHunt/db"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *db.Queries
	jwtSecretKey string
}

func main() {

	godotenv.Load()

	portS := os.Getenv("PORT")
	if portS == "" {
		fmt.Println("PORT is not set")
		return
	}
	fmt.Println("PORT is set to", portS)

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		fmt.Println("DB_URL is not set")
		return
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	conn, err := sql.Open("postgres", DB_URL)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("DB connection established")
	apiCfg := apiConfig{
		DB: db.New(conn),
		jwtSecretKey: jwtSecretKey,
	}
	fmt.Println(apiCfg)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1router := chi.NewRouter()
	v1router.Get("/healthz", handleHealth)
	v1router.Post("/user", apiCfg.handleCreateUser)
	v1router.Get("/user", apiCfg.handleGetJwt)
	v1router.Get("/userdata",apiCfg.JWTMiddleware(http.HandlerFunc(apiCfg.handleUserData)) )
	v1router.Get("/users", apiCfg.handleGetAllUsers)
	v1router.Post("/update", apiCfg.JWTMiddleware(http.HandlerFunc(apiCfg.handleUpdateUser)))

	router.Mount("/v1", v1router)

	srv := &http.Server{
		Addr:    ":" + portS,
		Handler: router,
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server is running on port", portS)

}
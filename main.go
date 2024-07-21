package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/harryfung22/ScrapeHive/internal/databse"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *databse.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	db := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", db)
	if err != nil {
		log.Fatal("Connection to database refused: ", err)
	}

	querries := databse.New(conn)

	apiConf := apiConfig{
		DB: querries,
	}

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	rss := chi.NewRouter()
	rss.Get("/health", handleRes)
	rss.Get("/err", handleErr)
	rss.Post("/users", apiConf.handleCreateUser)

	router.Mount("/rss", rss)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on port %v", port)
	er := server.ListenAndServe()
	if er != nil {
		log.Fatal((er))
	}

}

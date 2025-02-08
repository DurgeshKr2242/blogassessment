package main

import (
	"log"
	"os"

	"github.com/DurgeshKr2242/blogassessment/config"
	"github.com/DurgeshKr2242/blogassessment/db"
	`github.com/DurgeshKr2242/blogassessment/domains`
	"github.com/DurgeshKr2242/blogassessment/handlers"
	"github.com/DurgeshKr2242/blogassessment/router"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Connect to Database
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer database.Close()

	// 4. Initialize Domains (database interactions)
	blogPostDomain := domains.NewBlogPostDomain(database)

	// 5. Initialize Handlers
	blogPostHandlers := handlers.NewBlogPostHandler(blogPostDomain)

	// 6. Setup Router
	r := router.SetupRoutes(blogPostHandlers)

	// 7. Start the Server
	serverAddr := ":" + cfg.ServerPort
	log.Printf("Server listening on %s\n", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
	
	os.Exit(0)
}

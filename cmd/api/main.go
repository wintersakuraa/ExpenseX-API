package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/wintersakuraa/expense-x-api/internal/api"
	"github.com/wintersakuraa/expense-x-api/internal/handlers"
	"github.com/wintersakuraa/expense-x-api/internal/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	client, err := clerk.NewClient(os.Getenv("CLERK_SECRET_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewPostgresStorage(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("DB Connection Failed: %s\n", err)
	}

	var version string
	if err := db.QueryRow("select version()").Scan(&version); err != nil {
		panic(err)
	}
	fmt.Printf("version=%s\n", version)

	handler := handlers.New(client)

	s := api.NewServer(os.Getenv("PORT"), handler.LoadRoutes())
	go func() {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	log.Print("Server Started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}
	if err := db.Close(); err != nil {
		log.Fatalf("DB Close Failed: %+v", err)
	}

	log.Print("Server Exited Properly")
}

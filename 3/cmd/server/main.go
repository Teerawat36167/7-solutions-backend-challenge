package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Teerawat36167/PieFireDire/internal/api"
)

func main() {
	httpAddr := flag.String("http-addr", ":8080", "HTTP server address")
	grpcAddr := flag.String("grpc-addr", ":8081", "gRPC server address")
	flag.Parse()

	grpcServer, err := api.StartGRPCServer(*grpcAddr)
	if err != nil {
		log.Fatalf("Failed to set up gRPC server: %v", err)
	}

	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	handler := api.NewHandler()
	router := api.SetupRouter(handler)
	httpServer := &http.Server{
		Addr:    *httpAddr,
		Handler: router,
	}

	go func() {
		log.Printf("Starting gRPC server on %s", *grpcAddr)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	go func() {
		log.Printf("Starting HTTP server on %s", *httpAddr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}

	grpcServer.GracefulStop()
	log.Println("Servers successfully shut down")
}

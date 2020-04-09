package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/y-shika/todo_golang_api/web/router"
)

func main() {
	// TODO: この辺りはherokuとかでも大丈夫か確認する
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ここはgorillamuxの`mux.NewRouter`では動かなかった
	r := http.NewServeMux()
	r.Handle("/", router.New())

	server := &http.Server{
		Addr:         host + ":" + port,
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go func() {
		log.Printf("Start server on %s:%s (pid: %d)", host, port, os.Getpid())
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln("Server closed with error:", err)
		}
	}()

	// Ctrl+Cで中断した時の処理
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	log.Printf("SIGNAL %d received, then shutting down...\n", <-quit)

	// 指定した秒数後にキャンセルイベントが呼ばれる
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown: ", err)
	}
	log.Println("Server shutdown")
}

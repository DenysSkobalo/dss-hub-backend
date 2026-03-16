package main

import (
        "context"
        "log"
        "net/http"
        "os"
        "os/signal"
        "syscall"
        "time"
		"github.com/DenysSkobalo/dss-hub-backend/internal/middlewares"
)

func main() {
        log.Println("HUB-ANA-API: Hello, DSSpace! Ready for analytics.")

        mux := http.NewServeMux()
        mux.HandleFunc("/health", middlewares.SecurityMiddlewares("HUB-ANA-API", func(w http.ResponseWriter, r *http.Request) {
                w.Header().Set("Content-Type", "application/json")
                w.Write([]byte(`{"status":"online", "service":"HUB-ANA-API"}`))
        }))

        srv := &http.Server{
                Addr:    ":8080", 
                Handler: mux,
				ReadTimeout: 5 * time.Second,
				WriteTimeout: 10 * time.Second,
				IdleTimeout: 120 * time.Second, 
        }

        go func() {
                log.Printf("API listening on %s", srv.Addr)
                if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
                        log.Fatalf("Fatal: %v", err)
                }
        }()

        // Graceful Shutdown
        stop := make(chan os.Signal, 1)
        signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
        <-stop

        log.Println("SIGTERM received. Cleaning up...")
        ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
        defer cancel()

        if err := srv.Shutdown(ctx); err != nil {
                log.Printf("Error during shutdown: %v", err)
        }
        log.Println("HUB-ANA-API stopped safely.")
}

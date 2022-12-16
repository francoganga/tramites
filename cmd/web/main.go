package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/francoganga/go_reno2/pkg/routes"
	"github.com/francoganga/go_reno2/pkg/services"

    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/pgx"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Start a new container
	c := services.NewContainer()

    

    driver, err := pgx.WithInstance(c.Database, &pgx.Config{})

    if err != nil {
        panic(err)
    }

    m, err := migrate.NewWithDatabaseInstance("file://migrations", "app", driver)

    if err != nil {
        panic(err)
    }

    err = m.Up()

    if err != nil {
        panic(err)
    }









	defer func() {
		if err := c.Shutdown(); err != nil {
			c.Web.Logger.Fatal(err)
		}
	}()

	// Build the router
	routes.BuildRouter(c)

	// Start the server
	go func() {
		srv := http.Server{
			Addr:         fmt.Sprintf("%s:%d", c.Config.HTTP.Hostname, c.Config.HTTP.Port),
			Handler:      c.Web,
			ReadTimeout:  c.Config.HTTP.ReadTimeout,
			WriteTimeout: c.Config.HTTP.WriteTimeout,
			IdleTimeout:  c.Config.HTTP.IdleTimeout,
		}

		if c.Config.HTTP.TLS.Enabled {
			certs, err := tls.LoadX509KeyPair(c.Config.HTTP.TLS.Certificate, c.Config.HTTP.TLS.Key)
			if err != nil {
				c.Web.Logger.Fatalf("cannot load TLS certificate: %v", err)
			}

			srv.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{certs},
			}
		}

		if err := c.Web.StartServer(&srv); err != http.ErrServerClosed {
			c.Web.Logger.Fatalf("shutting down the server: %v", err)
		}
	}()

	// Start the scheduler service to queue periodic tasks
	go func() {
		if err := c.Tasks.StartScheduler(); err != nil {
			c.Web.Logger.Fatalf("scheduler shutdown: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, os.Kill)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := c.Web.Shutdown(ctx); err != nil {
		c.Web.Logger.Fatal(err)
	}
}

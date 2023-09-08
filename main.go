package main

import (
	"go-backend/conf"
	"go-backend/db"
	"go-backend/handlers"
	"go-backend/migrate"
	"go-backend/routes"
	"log"
	"os"
)

var (
	lstd *log.Logger
	lerr *log.Logger
)

// @title go-backend
// @version 1.0
// @description demo for mind
// @host localhost:3000
// @BasePath /
func main() {
	// Create a standard and the error logger objects.
	lstd = log.New(os.Stdout, "INFO:", log.LstdFlags|log.Lshortfile)
	lstd.Printf("Standard logger object %s created.", lstd.Prefix())
	lerr = log.New(os.Stdout, "ERROR:", log.LstdFlags|log.Lshortfile)
	lstd.Printf("Error logger object %s created.", lerr.Prefix())

	// Check if the file exists, and read the contents, unmarshal the yaml.
	c := &conf.Config{}
	if err := conf.CreateConfig(c); err != nil {
		lerr.Printf("Unable to create the config, %v.", err)
	}

	// Validate the configuratin attributes.
	if err := c.Validate(); err != nil {
		lerr.Printf("Unable to validate the config, %v.", err)
	}

	// Migarate database
	if err := migrate.MigrateDatabase(c); err != nil {
		lerr.Printf("Unable to migrate the database, %v.", err)
	} else {
		lstd.Printf("Migrate database completed.")
	}

	// The returned DB is safe for concurrent use by multiple goroutines
	// and maintains its own pool of idle connections. Thus, the OpenDB
	// function should be called just once. It is rarely necessary to
	// close a DB.
	pgsql, err := db.CreateSqlDB(c)
	if err != nil {
		lstd.Printf("Unable to create the database object, %v.", err)
		os.Exit(1)
	}

	lstd.Println("Connected to the database.")

	// Creating handlers for fiber route
	handlers := handlers.NewHandlers(lstd, lerr, c, pgsql)

	// Creating fiber
	routes := routes.Routes(handlers)

	// Fiber listen port
	routes.Listen(":3000")
}

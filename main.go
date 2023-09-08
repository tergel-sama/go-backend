package main

import (
	"go-backend/conf"
	"go-backend/db"
	"go-backend/handlers"
	"go-backend/migrate"
	"go-backend/routes"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/slog"
)

func main() {
	// creating customer log handler for slog
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Key = "date"

				outputLayout := "2006/01/02 15:04:05"
				formattedTimestamp := time.Now().Format(outputLayout)

				a.Value = slog.StringValue(formattedTimestamp)
			} else if a.Key == slog.SourceKey {
				a.Key = "file"

				if len(strings.Split(a.Value.String(), " ")) < 3 {
					return a
				}

				filepath := strings.Split(a.Value.String(), " ")[1]
				fileline := strings.Split(a.Value.String(), " ")[2]

				lastIndex := strings.LastIndex(filepath, "go-backend/")
				if lastIndex != -1 {
					msg := filepath[lastIndex+1+len("go-backend"):] + ":" + fileline[:len(fileline)-1]

					a.Value = slog.StringValue(msg)
				} else {
					a.Value = slog.StringValue("Not from our project")
				}
			}

			return a
		},
	})

	logger := slog.New(logHandler)

	// set customer logger to defualt slog
	slog.SetDefault(logger)

	// Check if the file exists, and read the contents, unmarshal the yaml.
	c := &conf.Config{}
	if err := conf.CreateConfig(c); err != nil {
		slog.Error("Unable to create the config", slog.Any("err", err))
		os.Exit(1)
	}

	// Validate the configuratin attributes.
	if err := c.Validate(); err != nil {
		slog.Error("Unable to validate the config", slog.Any("err", err))
		os.Exit(1)
	}

	// Migarate database
	if err := migrate.MigrateDatabase(c); err != nil {
		slog.Error("Unable to migrate the database", slog.Any("err", err))
		os.Exit(1)
	} else {
		slog.Info("Migrate database completed.")
	}

	// The returned DB is safe for concurrent use by multiple goroutines
	// and maintains its own pool of idle connections. Thus, the OpenDB
	// function should be called just once. It is rarely necessary to
	// close a DB.
	pgsql, err := db.CreateSqlDB(c)
	if err != nil {
		slog.Error("Unable to create the database object", slog.Any("err", err))
		os.Exit(1)
	}

	slog.Info("Connected to the database.")

	// Creating handlers for fiber route
	handlers := handlers.NewHandlers(c, pgsql)

	// Creating fiber
	routes := routes.Routes(handlers)

	// Fiber listen port
	routes.Listen(":3000")
}

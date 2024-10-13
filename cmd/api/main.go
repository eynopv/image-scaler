package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/eynopv/image-scaler/pkg/storage"
)

type config struct {
	port            int
	uploadDirectory string
}

type application struct {
	config  config
	logger  *slog.Logger
	storage *storage.FileSystemStorage
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 3000, "server port")
	flag.StringVar(&cfg.uploadDirectory, "dir", "./uploads", "upload directory")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		config:  cfg,
		logger:  logger,
		storage: storage.NewFileSystemStorage(cfg.uploadDirectory),
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", app.config.port),
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	if _, err := os.Stat(app.config.uploadDirectory); os.IsNotExist(err) {
		err := os.Mkdir(app.config.uploadDirectory, os.ModePerm)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}

	logger.Info("starting server", "address", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		logger.Error(err.Error())
	}
	os.Exit(1)
}

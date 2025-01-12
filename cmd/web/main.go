package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/Sa-YoorHeadley/snippetbox/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// Application Struct

type application struct {
    logger *slog.Logger
    snippets *models.SnippetModel
    templateCache map[string]*template.Template
}
//Main
func main() {
    
    
    // Flags
    addr := flag.String("addr", "localhost:4000", "HTTP network address")
    dsn := flag.String("dsn", "go:password@/snippetbox?parseTime=true", "MySQL data source name")

    flag.Parse()

    
    // Slog Logger
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    // loggerJSON := slog.New(slog.NewJSONHandler((os.Stdout, nil)))

    // Database connection

    db, err := openDB(*dsn)

    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }
    
    defer db.Close()

    templateCache, err := newTemplateCache()
    if err != nil {
        logger.Error(err.Error())
        os.Exit(1)
    }

    app := &application{
        logger: logger,
        snippets: &models.SnippetModel{DB: db},
        templateCache: templateCache,
    }

    logger.Info("Starting server", "addr", *addr)

    err = http.ListenAndServe(*addr, app.routes())

    logger.Error(err.Error())
    os.Exit(1)
}

func openDB(dsn string)(*sql.DB, error){
    db, err := sql.Open("mysql", dsn)

    if err != nil {
        return nil, err
    }

    err = db.Ping()

    if err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}
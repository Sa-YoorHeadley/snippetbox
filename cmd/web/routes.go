package main

import "net/http"

func (app *application) routes() *http.ServeMux {

   // Server Creation
    mux := http.NewServeMux()


    // Serving "static" files
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

    
    // Routes
    mux.HandleFunc("GET /{$}", app.home)
    mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
    mux.HandleFunc("GET /snippet/create", app.snippetCreate)
    mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
    
    return mux
}
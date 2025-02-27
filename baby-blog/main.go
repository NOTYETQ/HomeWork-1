// Filename: main.go

package main

import (
	"log"
	"net/http"
        "time"
)

func loggingMiddleware(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start:=time.Now()
            log.Printf("Timestamp: %s, Method: %s, URL: %s", start.Format(time.RFC3339), r.Method, r.URL.Path)
            next.ServeHTTP(w, r)
        })
     }

func homeHandler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./templates/home.html") 
}

func week1Handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./templates/week1.html")
}

func week2Handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./templates/week2.html")
}

func week3Handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./templates/week3.html")
}

func week4Handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./templates/week4.html")
}

func week5Handler(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./templates/week5.html")
}


func main() {
        mux := http.NewServeMux()

        // Serve the static files
        fileServer := http.FileServer(http.Dir("./static"))
        mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
        // Register the handlers
        mux.HandleFunc("/", homeHandler) 
        mux.HandleFunc("/week1", week1Handler)
        mux.HandleFunc("/week2", week2Handler)
        mux.HandleFunc("/week3", week3Handler)
        mux.HandleFunc("/week4", week4Handler)
        mux.HandleFunc("/week5", week5Handler)
        // Start listening for requests (start the web server)
        log.Println("Starting server on :4000")
        err := http.ListenAndServe(":4000", loggingMiddleware(mux))
	// Log error message if server quits unexpectedly
        if err != nil {
                log.Fatal(err)
        }
}
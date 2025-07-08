package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/goharbor/harbor/templates"
)

// Handler for serving static files (CSS, JS, images)
func staticHandler(w http.ResponseWriter, r *http.Request) {
	// Serve static files from src/portal/src
	staticPath := filepath.Join("src/portal/src", r.URL.Path)
	if _, err := os.Stat(staticPath); err == nil {
		http.ServeFile(w, r, staticPath)
		return
	}
	
	// Fallback to 404
	http.NotFound(w, r)
}

// Handler for the main dashboard
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// In a real implementation, you'd get this from session/auth
	props := templates.DashboardProps{
		IsSessionValid: true,
		AccountName:    "admin",
		CurrentLang:    "en",
	}
	
	ctx := context.Background()
	if err := templates.Dashboard(props).Render(ctx, w); err != nil {
		log.Printf("Error rendering dashboard: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Handler for login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Simple redirect to dashboard for now
	// In a real implementation, you'd render a login template
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	// Set up routes
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", loginHandler)
	
	// Serve static files
	http.HandleFunc("/images/", staticHandler)
	http.HandleFunc("/favicon.ico", staticHandler)
	
	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Starting templ server on port %s", port)
	log.Printf("Visit http://localhost:%s to view the new Harbor UI", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
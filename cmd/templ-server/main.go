package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

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

// Handler for projects page
func projectsHandler(w http.ResponseWriter, r *http.Request) {
	// Mock data - in real implementation, fetch from Harbor API
	projects := []templates.Project{
		{
			ID:          1,
			Name:        "library",
			Description: "Default public project for Harbor",
			Public:      true,
			RepoCount:   5,
			CreatedAt:   time.Now().AddDate(0, -2, 0),
			UpdatedAt:   time.Now().AddDate(0, 0, -1),
		},
		{
			ID:          2,
			Name:        "myproject",
			Description: "My private project for development",
			Public:      false,
			RepoCount:   12,
			CreatedAt:   time.Now().AddDate(0, -1, 0),
			UpdatedAt:   time.Now().AddDate(0, 0, -3),
		},
		{
			ID:          3,
			Name:        "production",
			Description: "Production container images",
			Public:      false,
			RepoCount:   8,
			CreatedAt:   time.Now().AddDate(0, -3, 0),
			UpdatedAt:   time.Now().AddDate(0, 0, -2),
		},
	}
	
	props := templates.ProjectsProps{
		IsSessionValid: true,
		AccountName:    "admin",
		CurrentLang:    "en",
		Projects:       projects,
		TotalCount:     len(projects),
	}
	
	ctx := context.Background()
	if err := templates.Projects(props).Render(ctx, w); err != nil {
		log.Printf("Error rendering projects: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Handler for login page
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Show login form
		props := templates.LoginProps{
			ErrorMessage: "",
			RedirectURL:  r.URL.Query().Get("redirect_url"),
		}
		
		ctx := context.Background()
		if err := templates.Login(props).Render(ctx, w); err != nil {
			log.Printf("Error rendering login: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		// Handle login form submission
		// In real implementation, validate credentials
		username := r.FormValue("username")
		password := r.FormValue("password")
		redirectURL := r.FormValue("redirect_url")
		
		if username == "admin" && password == "Harbor12345" {
			// Success - redirect to dashboard or original URL
			if redirectURL != "" {
				http.Redirect(w, r, redirectURL, http.StatusFound)
			} else {
				http.Redirect(w, r, "/", http.StatusFound)
			}
		} else {
			// Failed login
			props := templates.LoginProps{
				ErrorMessage: "Invalid username or password",
				RedirectURL:  redirectURL,
			}
			
			ctx := context.Background()
			if err := templates.Login(props).Render(ctx, w); err != nil {
				log.Printf("Error rendering login: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}
	}
}

// Handler for logout
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// In real implementation, clear session
	http.Redirect(w, r, "/login", http.StatusFound)
}

func main() {
	// Set up routes
	http.HandleFunc("/", dashboardHandler)
	http.HandleFunc("/dashboard", dashboardHandler)
	http.HandleFunc("/projects", projectsHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	
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
	log.Printf("Login with username: admin, password: Harbor12345")
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
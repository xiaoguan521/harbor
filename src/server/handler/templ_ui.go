package handler

import (
	"context"
	"net/http"
	"log"
	
	"github.com/goharbor/harbor/templates"
)

// TemplUIHandler handles requests for the templ-based UI
type TemplUIHandler struct {
	// Add any dependencies like session service, config, etc.
}

// NewTemplUIHandler creates a new templ UI handler
func NewTemplUIHandler() *TemplUIHandler {
	return &TemplUIHandler{}
}

// Dashboard serves the main dashboard page
func (h *TemplUIHandler) Dashboard(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user session and authentication status
	// For now, using mock data
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

// RegisterRoutes registers all templ UI routes with the router
func (h *TemplUIHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Dashboard)
	mux.HandleFunc("/dashboard", h.Dashboard)
	
	// Add more routes as needed
	// mux.HandleFunc("/projects", h.Projects)
	// mux.HandleFunc("/users", h.Users)
	// mux.HandleFunc("/logs", h.Logs)
}
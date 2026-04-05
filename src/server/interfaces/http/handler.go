package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"easy-login/server/application"
	"easy-login/server/domain"
)

type Handler struct {
	service        application.Service
	allowedOrigins map[string]struct{}
}

type createGuestIdentityRequest struct {
	DisplayName string `json:"display_name"`
}

type resumeIdentityRequest struct {
	DeviceToken string `json:"device_token"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func NewHandler(service application.Service) Handler {
	return Handler{
		service:        service,
		allowedOrigins: loadAllowedOrigins(),
	}
}

func (h Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /identities/guest", h.createGuestIdentity)
	mux.HandleFunc("POST /identities/resume", h.resumeIdentity)
	mux.HandleFunc("GET /healthz", h.healthz)
	mux.HandleFunc("GET /readyz", h.readyz)
	return h.withCORS(mux)
}

func (h Handler) createGuestIdentity(w http.ResponseWriter, r *http.Request) {
	var request createGuestIdentityRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid json body"})
		return
	}

	response, err := h.service.CreateGuestIdentity(r.Context(), application.CreateGuestIdentityInput{
		DisplayName: request.DisplayName,
	})
	if err != nil {
		if errors.Is(err, domain.ErrEmptyDisplayName) {
			writeJSON(w, http.StatusBadRequest, errorResponse{Error: err.Error()})
			return
		}

		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal error"})
		return
	}

	writeJSON(w, http.StatusCreated, response)
}

func (h Handler) resumeIdentity(w http.ResponseWriter, r *http.Request) {
	var request resumeIdentityRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid json body"})
		return
	}

	response, err := h.service.ResumeIdentityFromDeviceToken(r.Context(), request.DeviceToken)
	if err != nil {
		if errors.Is(err, application.ErrIdentityNotFound) {
			writeJSON(w, http.StatusNotFound, errorResponse{Error: "identity not found"})
			return
		}

		writeJSON(w, http.StatusInternalServerError, errorResponse{Error: "internal error"})
		return
	}

	writeJSON(w, http.StatusOK, response)
}

func (h Handler) healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) readyz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if _, ok := h.allowedOrigins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loadAllowedOrigins() map[string]struct{} {
	configured := os.Getenv("CORS_ALLOWED_ORIGINS")
	if strings.TrimSpace(configured) == "" {
		configured = strings.Join([]string{
			"http://localhost:5173",
			"http://127.0.0.1:5173",
			"http://localhost:4173",
			"http://127.0.0.1:4173",
		}, ",")
	}

	allowedOrigins := map[string]struct{}{}
	for _, origin := range strings.Split(configured, ",") {
		trimmed := strings.TrimSpace(origin)
		if trimmed == "" {
			continue
		}
		allowedOrigins[trimmed] = struct{}{}
	}

	return allowedOrigins
}

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

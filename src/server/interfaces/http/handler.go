package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"easy-login/server/application"
	"easy-login/server/domain"
)

type Handler struct {
	service application.Service
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
	return Handler{service: service}
}

func (h Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /identities/guest", h.createGuestIdentity)
	mux.HandleFunc("POST /identities/resume", h.resumeIdentity)
	mux.HandleFunc("GET /healthz", h.healthz)
	mux.HandleFunc("GET /readyz", h.readyz)
	return mux
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

func writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

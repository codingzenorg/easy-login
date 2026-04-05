package httpapi_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"easy-login/server/application"
	"easy-login/server/infrastructure/memory"
	httpapi "easy-login/server/interfaces/http"
)

type fixedPlayerIDGenerator struct {
	value string
}

func (g fixedPlayerIDGenerator) NewPlayerID(context.Context) (string, error) {
	return g.value, nil
}

type fixedDeviceTokenGenerator struct {
	value string
}

func (g fixedDeviceTokenGenerator) NewDeviceToken(context.Context) (string, error) {
	return g.value, nil
}

func testHandler() http.Handler {
	service := application.NewService(
		memory.NewPlayerRepository(),
		memory.NewDeviceRegistrationRepository(),
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	return httpapi.NewHandler(service).Routes()
}

func TestCreateGuestIdentityEndpoint(t *testing.T) {
	handler := testHandler()

	request := httptest.NewRequest(http.MethodPost, "/identities/guest", bytes.NewBufferString(`{"display_name":"henrique"}`))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected 201 Created, got %d", recorder.Code)
	}

	var body map[string]string
	if err := json.Unmarshal(recorder.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected json response: %v", err)
	}

	if body["player_id"] != "player-001" {
		t.Fatalf("expected stable player id, got %q", body["player_id"])
	}

	if body["device_token"] != "device-001" {
		t.Fatalf("expected device token, got %q", body["device_token"])
	}
}

func TestResumeIdentityEndpointReturnsNotFoundForUnknownToken(t *testing.T) {
	handler := testHandler()

	request := httptest.NewRequest(http.MethodPost, "/identities/resume", bytes.NewBufferString(`{"device_token":"missing"}`))
	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected 404 Not Found, got %d", recorder.Code)
	}
}

func TestHealthAndReadyEndpoints(t *testing.T) {
	handler := testHandler()

	for _, path := range []string{"/healthz", "/readyz"} {
		request := httptest.NewRequest(http.MethodGet, path, nil)
		recorder := httptest.NewRecorder()

		handler.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusOK {
			t.Fatalf("expected %s to return 200 OK, got %d", path, recorder.Code)
		}
	}
}

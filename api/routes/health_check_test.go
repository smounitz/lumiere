package routes

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestHealthCheck(t *testing.T){
	t.Run("return sttus ok", func(t *testing.T){
		request, _ := http.NewRequest(http.MethodGet, "v1/status", nil)
		response := httptest.NewRecorder()
		
		newHealthCheckRoute := NewHealthCheckRoute()
		newHealthCheckRoute(response, request, nil)

		got := response.Body.String()
		want := `"status":"ok"`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
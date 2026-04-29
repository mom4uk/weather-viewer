package testutils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func AssertStatus(t *testing.T, rr *httptest.ResponseRecorder, code int) {
	t.Helper()

	if rr.Code != code {
		t.Fatalf("expected %d, got %d\nbody: %s", code, rr.Code, rr.Body.String())
	}
}

func PerformRequest(t *testing.T, app *TestApp, method, path string, body io.Reader) *httptest.ResponseRecorder {
	t.Helper()

	req, err := http.NewRequest(method, path, body)
	require.NoError(t, err)

	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	rr := httptest.NewRecorder()
	app.Server.Mux.ServeHTTP(rr, req)

	return rr
}

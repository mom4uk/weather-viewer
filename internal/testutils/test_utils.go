package testutils

import (
	"net/http/httptest"
	"testing"
)

func AssertStatus(t *testing.T, rr *httptest.ResponseRecorder, code int) {
	t.Helper()

	if rr.Code != code {
		t.Fatalf("expected %d, got %d\nbody: %s", code, rr.Code, rr.Body.String())
	}
}

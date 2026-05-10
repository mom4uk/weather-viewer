package testutils

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"weather-viewer/internal/interfaces"

	"github.com/stretchr/testify/require"
)

func AssertStatus(t *testing.T, rr *httptest.ResponseRecorder, code int) {
	t.Helper()

	if rr.Code != code {
		t.Fatalf("expected %d, got %d\nbody: %s", code, rr.Code, rr.Body.String())
	}
}

func PerformRequest(t *testing.T, app *TestApp, method, path string, body io.Reader, sessionToken string) *httptest.ResponseRecorder {
	t.Helper()

	req, err := http.NewRequest(method, path, body)
	require.NoError(t, err)

	const SessionCookieName = "session_token"
	req.AddCookie(&http.Cookie{
		Name:  SessionCookieName,
		Value: sessionToken,
	})

	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	rr := httptest.NewRecorder()
	app.Server.Mux.ServeHTTP(rr, req)

	return rr
}

// тут похоже проблема с архитектурой. Из-за того, что я передаю в NewLocationService() weatherClient, я вынужден
// либо кидать его в каждый SetupTests, либо подменять на FakeWeatherClient. Мне кажется это не правильно.
func SetupTests(t *testing.T) (*TestApp, *TestDB) {
	db := NewTestDB()
	app := NewTestApp(db)

	err := TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")
	return app, db
}

func SetupTestWithWeather(t *testing.T, weatherClient interfaces.Weather) (*TestApp, *TestDB) {
	db := NewTestDB()
	app := NewTestAppForWeather(db, weatherClient)

	err := TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")
	return app, db
}

package testutils

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const TestLogin = "test1234"
const TestPassword = "qwerty1234"

func SignInUser(t *testing.T, app *TestApp, login, password string) string {
	form := url.Values{}
	form.Set("login", login)
	form.Set("password", password)

	rr := PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/login",
		strings.NewReader(form.Encode()),
		"",
	)
	AssertStatus(t, rr, http.StatusOK)

	cookies := rr.Header()["Set-Cookie"]

	var sessionToken string
	for _, c := range cookies {
		if strings.Contains(c, "session_token=") {
			sessionToken = strings.Split(strings.Split(c, "=")[1], ";")[0]
		}
	}

	return sessionToken
}

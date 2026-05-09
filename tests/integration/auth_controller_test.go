package integration

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/dto"
	"weather-viewer/internal/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// POST /auth/register

func TestRegistration_success(t *testing.T) {
	app, _ := testutils.SetupTests(t)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/register",
		strings.NewReader("login=loginLength20simbols&password=passwLength20simbols"),
		"",
	)

	testutils.AssertStatus(t, rr, http.StatusCreated)
	var got dto.UserResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := dto.UserResponse{
		Login: "loginLength20simbols",
	}

	assert.Equal(t, expected, got)
}

func TestRegistration_success_loginWithSpaces(t *testing.T) {
	app, _ := testutils.SetupTests(t)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/register",
		strings.NewReader("login=  loginLength20simbols   &password=6chars"),
		"",
	)

	testutils.AssertStatus(t, rr, http.StatusCreated)
	var got dto.UserResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := dto.UserResponse{
		Login: "loginLength20simbols",
	}

	assert.Equal(t, expected, got)
}

func TestRegistration_error_invalidLogin(t *testing.T) {
	app, _ := testutils.SetupTests(t)

	loginInvalidLengthMessage := "Длина логина должен быть от 6 до 20 символов"
	invalidLoginMessage := "Логин может состоять только из латинских букв и цифр"
	testData := []struct {
		name    string
		login   string
		message string
	}{
		{"Dots are not allowed", "login=testTest.&password=password", invalidLoginMessage},
		{"Special chars are not allowed", "login=2*()_-?!test&password=password", invalidLoginMessage},
		{"Login over 20 simbols are not allowed", "login=loginLength21simbolss&password=password", loginInvalidLengthMessage},
		{"Login less than 6 simbols are not allowed", "login=less&password=password", loginInvalidLengthMessage},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {

			rr := testutils.PerformRequest(
				t,
				app,
				http.MethodPost,
				"/auth/register",
				strings.NewReader(tt.login),
				"",
			)
			testutils.AssertStatus(t, rr, http.StatusUnprocessableEntity)

			var got domain.ErrorResponse
			require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

			expected := domain.ErrorResponse{
				Message: tt.message,
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestRegistration_error_nonUniqueLogin(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/register",
		strings.NewReader("login=test1234&password=password"),
		"",
	)

	testutils.AssertStatus(t, rr, http.StatusConflict)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Пользователь с таким логином уже существует",
	}

	assert.Equal(t, expected, got)
}

func TestRegistration_error_invalidPassword(t *testing.T) {
	app, _ := testutils.SetupTests(t)
	errorMessage := "Длина пароля должна быть от 6 до 20 символов"

	testData := []struct {
		name    string
		login   string
		message string
	}{

		{"Password over 20 simbols are not allowed", "login=test1111&password=passwoLength21simbols", errorMessage},
		{"Password less than 6 simbols are not allowed", "login=test1111&password=5char", errorMessage},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			rr := testutils.PerformRequest(
				t,
				app,
				http.MethodPost,
				"/auth/register",
				strings.NewReader(tt.login),
				"",
			)
			testutils.AssertStatus(t, rr, http.StatusUnprocessableEntity)

			var got domain.ErrorResponse
			require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

			expected := domain.ErrorResponse{
				Message: tt.message,
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestRegistration_error_absenceOfFields(t *testing.T) {
	app, _ := testutils.SetupTests(t)
	errorMessage := "Не передан логин и/или пароль"

	testData := []struct {
		name    string
		login   string
		message string
	}{

		{"Absence of login and password is not allowed", "", errorMessage},
		{"Empty password is not allowed", "login=test1234&password=", errorMessage},
		{"Empty login is not allowed", "login=&password=test1234", errorMessage},
		{"Absence of password is not allowed", "login=loginLength20simbols", errorMessage},
		{"Absence of login is not allowed", "password=loginLength20simbols", errorMessage},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			rr := testutils.PerformRequest(
				t,
				app,
				http.MethodPost,
				"/auth/register",
				strings.NewReader(tt.login),
				"",
			)
			testutils.AssertStatus(t, rr, http.StatusBadRequest)

			var got domain.ErrorResponse
			require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

			expected := domain.ErrorResponse{
				Message: tt.message,
			}

			assert.Equal(t, expected, got)
		})
	}
}

// POST /auth/login

func TestLogin_success(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/login",
		strings.NewReader("login=test1234&password=qwerty1234"),
		"",
	)
	testutils.AssertStatus(t, rr, http.StatusOK)
}

func TestLogin_error_incorrectLogin(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/login",
		strings.NewReader("login=incorrectLogin1&password=qwerty1234"),
		"",
	)
	testutils.AssertStatus(t, rr, http.StatusUnauthorized)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))
	expected := domain.ErrorResponse{
		Message: "Неверный логин или пароль",
	}
	assert.Equal(t, expected, got)
}

func TestLogin_error_incorrectPassword(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/login",
		strings.NewReader("login=test1234&password=incorrectPassword"),
		"",
	)
	testutils.AssertStatus(t, rr, http.StatusUnauthorized)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))
	expected := domain.ErrorResponse{
		Message: "Неверный логин или пароль",
	}
	assert.Equal(t, expected, got)
}

func TestLogin_error_absenceOfFields(t *testing.T) {
	app, _ := testutils.SetupTests(t)
	errorMessage := "Не передан логин и/или пароль"

	testData := []struct {
		name    string
		login   string
		message string
	}{

		{"Absence of login and password is not allowed", "", errorMessage},
		{"Empty password is not allowed", "login=test1234&password=", errorMessage},
		{"Empty login is not allowed", "login=&password=test1234", errorMessage},
		{"Absence of password is not allowed", "login=loginLength20simbols", errorMessage},
		{"Absence of login is not allowed", "password=loginLength20simbols", errorMessage},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			rr := testutils.PerformRequest(
				t,
				app,
				http.MethodPost,
				"/auth/register",
				strings.NewReader(tt.login),
				"",
			)
			testutils.AssertStatus(t, rr, http.StatusBadRequest)

			var got domain.ErrorResponse
			require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

			expected := domain.ErrorResponse{
				Message: tt.message,
			}

			assert.Equal(t, expected, got)
		})
	}
}

// POST /auth/logout

func TestLogout_success(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/logout",
		strings.NewReader("login=test1234"),
		sessionID,
	)
	testutils.AssertStatus(t, rr, http.StatusNoContent)
}

func TestLogout_success_incorrectSessionID(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/auth/logout",
		strings.NewReader("login=test1234"),
		"1",
	)
	testutils.AssertStatus(t, rr, http.StatusNoContent)
}

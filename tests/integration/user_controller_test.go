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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	rr := testutils.PerformRequest( // подумать над тем как вынести session_token отсюда, он не всегда нужен
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	rr := testutils.PerformRequest( // подумать над тем как вынести session_token отсюда, он не всегда нужен
		t,
		app,
		http.MethodPost,
		"/auth/register",
		strings.NewReader("login=loginLength20simbols   &password=6chars"),
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
	loginInvalidLengthMessage := "Длина логина должен быть от 6 до 20 символов"
	invalidLoginMessage := "Логин может состоять только из латинских букв и цифр"
	testData := []struct {
		name    string
		login   string
		message string
	}{
		{"Dots are not allowed", "login=testTest.", invalidLoginMessage},
		{"Special chars are not allowed", "login=2*()_-?!test", invalidLoginMessage},
		{"Empty login is not allowed", "login=", loginInvalidLengthMessage},
		{"Login over 20 simbols are not allowed", "login=loginLength21simbolss", loginInvalidLengthMessage},
		{"Login less than 6 simbols are not allowed", "login=less", loginInvalidLengthMessage},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			db := testutils.NewTestDB()
			app := testutils.NewTestApp(db)

			err := testutils.TruncateAll(db.DB)
			require.NoError(t, err, "truncate error")

			rr := testutils.PerformRequest( // подумать над тем как вынести session_token отсюда, он не всегда нужен
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	rr := testutils.PerformRequest( // подумать над тем как вынести session_token отсюда, он не всегда нужен
		t,
		app,
		http.MethodPost,
		"/auth/register",
		strings.NewReader("login=test1234"),
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
	errorMessage := "Длина пароля должна быть от 6 до 20 символов"

	testData := []struct {
		name    string
		login   string
		message string
	}{

		{"Empty password is not allowed", "login=test1111&password=", errorMessage},
		{"Password over 20 simbols are not allowed", "login=test1111&password=passwoLength21simbols", errorMessage},
		{"Password less than 6 simbols are not allowed", "login=test1111&password=5char", errorMessage},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			db := testutils.NewTestDB()
			app := testutils.NewTestApp(db)

			err := testutils.TruncateAll(db.DB)
			require.NoError(t, err, "truncate error")

			rr := testutils.PerformRequest( // подумать над тем как вынести session_token отсюда, он не всегда нужен
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
	errorMessage := "Не передан логин и/или пароль"

	testData := []struct {
		name    string
		login   string
		message string
	}{

		{"Empty password is not allowed", "", errorMessage},
		{"Password over 20 simbols are not allowed", "login=loginLength20simbols", errorMessage},
		{"Password less than 6 simbols are not allowed", "password=loginLength20simbols", errorMessage},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			db := testutils.NewTestDB()
			app := testutils.NewTestApp(db)

			err := testutils.TruncateAll(db.DB)
			require.NoError(t, err, "truncate error")

			rr := testutils.PerformRequest( // подумать над тем как вынести session_token отсюда, он не всегда нужен
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

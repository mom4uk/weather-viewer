package integration

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"weather-viewer/internal/clients"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/dto"
	"weather-viewer/tests/fixtures"
	"weather-viewer/tests/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sessionID = "550e8400-e29b-41d4-a716-446655440000"

// GET /searchLocation/{id}
func TestSearchLocation_success(t *testing.T) {
	server := testutils.NewErrorServer(200)
	defer server.Close()

	weatherClient := &clients.WeatherClient{
		URL:    server.URL,
		APIKey: "key",
		Client: server.Client(),
	}

	db := testutils.NewTestDB()
	app := testutils.NewTestAppForWeather(db, weatherClient)

	err := testutils.TruncateAll(db.Postgres)
	require.NoError(t, err, "truncate db error")

	err = testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedLocations(db.Postgres)
	require.NoError(t, err, "seed locations error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/searchLocation/1",
		nil,
		token,
	)

	testutils.AssertStatus(t, rr, http.StatusOK)

	var got dto.LocationResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := dto.LocationResponse{
		ID:        1,
		Name:      "Москва",
		UserID:    1,
		Latitude:  55.7522,
		Longitude: 37.6156,
		Weather:   fixtures.GetMoscowWeather(),
	}

	assert.Equal(t, expected, got)
}

func TestSearchLocation_weatherAPIErrors(t *testing.T) {

	tests := []struct {
		name           string
		weatherStatus  int
		expectedStatus int
	}{
		{
			name:           "weather api 404",
			weatherStatus:  http.StatusNotFound,
			expectedStatus: http.StatusBadGateway,
		},
		{
			name:           "weather api 500",
			weatherStatus:  http.StatusInternalServerError,
			expectedStatus: http.StatusBadGateway,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := testutils.NewErrorServer(tt.weatherStatus)
			defer server.Close()

			weatherClient := &clients.WeatherClient{
				URL:    server.URL,
				APIKey: "key",
				Client: server.Client(),
			}

			db := testutils.NewTestDB()
			app := testutils.NewTestAppForWeather(db, weatherClient)

			token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

			err := testutils.TruncateAll(db.Postgres)
			require.NoError(t, err, "truncate error")

			err = testutils.SeedUsers(db.Postgres)
			require.NoError(t, err, "seed users error")

			err = testutils.SeedLocations(db.Postgres)
			require.NoError(t, err, "seed locations error")

			rr := testutils.PerformRequest(
				t,
				app,
				http.MethodGet,
				"/searchLocation/1",
				nil,
				token,
			)

			testutils.AssertStatus(t, rr, tt.expectedStatus)

			var got domain.ErrorResponse
			require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

			assert.NotEmpty(t, got.Message)
		})
	}
}

func TestSearchLocation_error_incorrectId(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/searchLocation/aaa",
		nil,
		token,
	)

	testutils.AssertStatus(t, rr, http.StatusBadRequest)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Некорректное значение в id",
	}

	assert.Equal(t, expected, got)
}

func TestSearchLocation_error_locationNotFound(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/searchLocation/244",
		nil,
		token,
	)

	testutils.AssertStatus(t, rr, http.StatusNotFound)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Данная локация не найдена",
	}

	assert.Equal(t, expected, got)
}

// POST /addLocation

func TestAddLocation_success(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/addLocation",
		strings.NewReader("name=Тверь&id=1&latitude=3&longitude=4"),
		token,
	)

	testutils.AssertStatus(t, rr, http.StatusCreated)

	var got dto.LocationResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	assert.NotZero(t, got.ID)

	expected := dto.LocationResponse{
		ID: 0, Name: "Тверь", UserID: 1, Latitude: 3, Longitude: 4,
	}
	got.ID = 0
	assert.Equal(t, expected, got)
}

func TestAddLocation_error_invalidFieldValues(t *testing.T) {
	app, db := testutils.SetupTests(t)
	nameMessage := "Некорректное значение в name"
	latitudeMessage := "Некорректное значение в latitude"
	longitudeMessage := "Некорректное значение в longitude"

	tests := []struct {
		name    string
		input   string
		message string
	}{
		{"empty name", "name=&latitude=3&longitude=4", nameMessage},
		{"name with spaces for name", "name=  &latitude=3&longitude=4", nameMessage},
		{"name with special char for name", "name=...&latitude=3&longitude=4", nameMessage},
		{"name with special char for name", "name=!?л?&latitude=3&longitude=4", nameMessage},
		{"name with special char for name", "name=^^k^&latitude=3&longitude=4", nameMessage},
		{"name with numbers for name", "name=1234d5&latitude=3&longitude=4", nameMessage},

		{"empty latitude", "name=Тверь&latitude=&longitude=4", latitudeMessage},
		{"latitude with spaces", "name=Тверь&latitude=   &longitude=4", latitudeMessage},
		{"latitude with cyrillic char", "name=Тверь&latitude=лво12&longitude=4", latitudeMessage},
		{"latitude with special char", "name=Тверь&latitude=:3:&longitude=4", latitudeMessage},
		{"latitude with english char", "name=Тверь&latitude=abc2&longitude=4", latitudeMessage},

		{"empty longitude", "name=Тверь&latitude=3&longitude=", longitudeMessage},
		{"longitude with spaces", "name=Тверь&latitude=3&longitude=    ", longitudeMessage},
		{"longitude with cyrillic char", "name=Тверь&latitude=3&longitude=ю.12", longitudeMessage},
		{"longitude with english chars", "name=Тверь&latitude=3&longitude=test", longitudeMessage},
		{"longitude with special char", "name=Тверь&latitude=3&longitude=№№2", longitudeMessage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testutils.TruncateAll(db.Postgres)
			require.NoError(t, err, "truncate error")

			err = testutils.SeedUsers(db.Postgres)
			require.NoError(t, err, "seed users error")

			token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

			rr := testutils.PerformRequest(
				t,
				app,
				http.MethodPost,
				"/addLocation",
				strings.NewReader(tt.input),
				token,
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

func TestAddLocation_error_locationAlreadyExists(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedLocations(db.Postgres)
	require.NoError(t, err, "seed locations error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/addLocation",
		strings.NewReader("name=Москва&latitude=0&longitude=0"),
		token,
	)

	testutils.AssertStatus(t, rr, http.StatusConflict)
	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Такая локация уже существует",
	}

	assert.Equal(t, expected, got)
}

// Auth test
func TestAuth_error_absenceOfSessionId(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedLocations(db.Postgres)
	require.NoError(t, err, "seed locations error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/addLocation",
		strings.NewReader("name=Москва&latitude=0&longitude=0"),
		"",
	)

	testutils.AssertStatus(t, rr, http.StatusUnauthorized)
	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Unauthorized",
	}

	assert.Equal(t, expected, got)
}

// GET /getLocations

func TestGetLocations_success(t *testing.T) {
	server := testutils.NewErrorServer(200)
	defer server.Close()

	weatherClient := &clients.WeatherClient{
		URL:    server.URL,
		APIKey: "key",
		Client: server.Client(),
	}

	db := testutils.NewTestDB()
	app := testutils.NewTestAppForWeather(db, weatherClient)

	err := testutils.TruncateAll(db.Postgres)
	require.NoError(t, err, "truncate db error")

	err = testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedLocations(db.Postgres)
	require.NoError(t, err, "seed locations error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/getLocations",
		nil,
		token,
	)

	testutils.AssertStatus(t, rr, http.StatusOK)

	var got []dto.LocationResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := []dto.LocationResponse{
		{
			ID:        1,
			Name:      "Москва",
			UserID:    1,
			Latitude:  55.7522,
			Longitude: 37.6156,
			Weather:   fixtures.GetMoscowWeather(),
		},
		{
			ID:        2,
			Name:      "Санкт-Петербург",
			UserID:    1,
			Latitude:  59.8944,
			Longitude: 30.2642,
			Weather:   fixtures.GetSpbWeather(),
		},
	}

	assert.Equal(t, expected, got)
}

// DELETE /removeLocation/{id}

func TestRemoveLocation_success(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedLocations(db.Postgres)
	require.NoError(t, err, "seed locations error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodDelete,
		"/removeLocation/1",
		nil,
		token,
	)
	testutils.AssertStatus(t, rr, http.StatusNoContent)

	rr = testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/searchLocation/1",
		nil,
		token,
	)

	testutils.AssertStatus(t, rr, http.StatusNotFound)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Данная локация не найдена",
	}

	assert.Equal(t, expected, got)
}

// можно накинуть проверок через test table
func TestRemoveLocation_error_invalidId(t *testing.T) {
	app, db := testutils.SetupTests(t)

	err := testutils.SeedUsers(db.Postgres)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedLocations(db.Postgres)
	require.NoError(t, err, "seed locations error")

	token := testutils.SignInUser(t, app, testutils.TestLogin, testutils.TestPassword)

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodDelete,
		"/removeLocation/1a4",
		nil,
		token,
	)
	testutils.AssertStatus(t, rr, http.StatusBadRequest)

	var got domain.ErrorResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := domain.ErrorResponse{
		Message: "Некорректное значение в id",
	}

	assert.Equal(t, expected, got)
}

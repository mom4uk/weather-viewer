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

const sessionID = "550e8400-e29b-41d4-a716-446655440000"

// GET /searchLocation/{id}
func TestSearchLocation_success(t *testing.T) {
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	err = testutils.SeedLocations(db.DB)
	require.NoError(t, err, "seed locations error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/searchLocation/1",
		nil,
		sessionID,
	)

	testutils.AssertStatus(t, rr, http.StatusOK)

	var got dto.LocationResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := dto.LocationResponse{
		ID: 1, Name: "Москва", UserID: 1, Latitude: 0, Longitude: 0,
	}

	assert.Equal(t, expected, got)
}

func TestSearchLocation_error_incorrectId(t *testing.T) {
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/searchLocation/aaa",
		nil,
		sessionID,
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/searchLocation/244",
		nil,
		sessionID,
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/addLocation",
		strings.NewReader("name=Тверь&id=1&latitude=3&longitude=4"),
		sessionID,
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	nameMessage := "Некорректное значение в name"
	idMessage := "Некорректное значение в id"
	latitudeMessage := "Некорректное значение в latitude"
	longitudeMessage := "Некорректное значение в longitude"

	tests := []struct {
		name    string
		input   string
		message string
	}{
		{"empty name", "name=&id=1&latitude=3&longitude=4", nameMessage},
		{"name with spaces for name", "name=  &id=1&latitude=3&longitude=4", nameMessage},
		{"name with special char for name", "name=...&id=1&latitude=3&longitude=4", nameMessage},
		{"name with special char for name", "name=!?л?&id=1&latitude=3&longitude=4", nameMessage},
		{"name with special char for name", "name=^^k^&id=1&latitude=3&longitude=4", nameMessage},
		{"name with numbers for name", "name=1234d5&id=1&latitude=3&longitude=4", nameMessage},

		{"empty id", "name=Тверь&id=&latitude=3&longitude=4", idMessage},
		{"id with spaces", "name=Тверь&id= &latitude=3&longitude=4", idMessage},
		{"id with cyrillic char", "name=Тверь&id=тест&latitude=3&longitude=4", idMessage},
		{"id with special char", "name=Тверь&id=**3*&latitude=3&longitude=4", idMessage},
		{"id with english char", "name=Тверь&id=test1&latitude=3&longitude=4", idMessage},

		{"empty latitude", "name=Тверь&id=1&latitude=&longitude=4", latitudeMessage},
		{"latitude with spaces", "name=Тверь&id=1&latitude=   &longitude=4", latitudeMessage},
		{"latitude with cyrillic char", "name=Тверь&id=1&latitude=лво12&longitude=4", latitudeMessage},
		{"latitude with special char", "name=Тверь&id=1&latitude=:3:&longitude=4", latitudeMessage},
		{"latitude with english char", "name=Тверь&id=1&latitude=abc2&longitude=4", latitudeMessage},

		{"empty longitude", "name=Тверь&id=1&latitude=3&longitude=", longitudeMessage},
		{"longitude with spaces", "name=Тверь&id=1&latitude=3&longitude=    ", longitudeMessage},
		{"longitude with cyrillic char", "name=Тверь&id=1&latitude=3&longitude=ю.12", longitudeMessage},
		{"longitude with english chars", "name=Тверь&id=1&latitude=3&longitude=test", longitudeMessage},
		{"longitude with special char", "name=Тверь&id=1&latitude=3&longitude=№№2", longitudeMessage},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := testutils.TruncateAll(db.DB)
			require.NoError(t, err, "truncate error")

			err = testutils.SeedUsers(db.DB)
			require.NoError(t, err, "seed users error")

			err = testutils.SeedSession(db.DB, sessionID)
			require.NoError(t, err, "seed sessions error")

			rr := testutils.PerformRequest(
				t,
				app,
				http.MethodPost,
				"/addLocation",
				strings.NewReader(tt.input),
				sessionID,
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	err = testutils.SeedLocations(db.DB)
	require.NoError(t, err, "seed locations error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/addLocation",
		strings.NewReader("name=Москва&id=1&latitude=0&longitude=0"),
		sessionID,
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")

	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")

	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")

	err = testutils.SeedLocations(db.DB)
	require.NoError(t, err, "seed locations error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodPost,
		"/addLocation",
		strings.NewReader("name=Москва&id=1&latitude=0&longitude=0"),
		"test-session",
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
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	err := testutils.TruncateAll(db.DB)
	require.NoError(t, err, "truncate error")
	err = testutils.SeedUsers(db.DB)
	require.NoError(t, err, "seed users error")
	err = testutils.SeedSession(db.DB, sessionID)
	require.NoError(t, err, "seed sessions error")
	err = testutils.SeedLocations(db.DB)
	require.NoError(t, err, "seed locations error")

	rr := testutils.PerformRequest(
		t,
		app,
		http.MethodGet,
		"/getLocations",
		nil,
		sessionID,
	)

	testutils.AssertStatus(t, rr, http.StatusOK)

	var got []dto.LocationResponse
	require.NoError(t, json.NewDecoder(rr.Body).Decode(&got))

	expected := []dto.LocationResponse{
		{ID: 1, Name: "Москва", UserID: 1, Latitude: 0, Longitude: 0},
		{ID: 2, Name: "Санкт-Петербург", UserID: 1, Latitude: 1, Longitude: 1},
	}

	assert.Equal(t, expected, got)
}

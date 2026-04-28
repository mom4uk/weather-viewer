package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"weather-viewer/internal/domain"
	"weather-viewer/internal/dto"
	testutils2 "weather-viewer/internal/testutils"
	"weather-viewer/tests/seeds"
)

func TestSearchLocation_success(t *testing.T) {
	db := testutils2.NewTestDB()
	app := testutils2.NewTestApp(db)

	if err := seeds.AddUser(db); err != nil {
		t.Fatalf("data seed error %s", err)
	}

	if err := seeds.AddLocations(db); err != nil {
		t.Fatalf("data seed error %s", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/searchLocation/1", nil)
	if err != nil {
		t.Fatalf("http request error: %s", err)
	}
	rr := httptest.NewRecorder()

	app.Server.Mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d\nbody: %s", rr.Code, rr.Body.String())
	}
	var got []dto.LocationResponse
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("faild to decode response: %v", err)
	}

	expected := dto.LocationResponse{
		ID: 1, Name: "Москва", UserID: 1, Latitude: 0, Longitude: 0,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %+v, got %+v", expected, got)
	}
}

func TestSearchLocation_error_incorrectId(t *testing.T) {
	db := testutils2.NewTestDB()
	app := testutils2.NewTestApp(db)

	req, err := http.NewRequest(http.MethodGet, "/searchLocation/aaa", nil)
	if err != nil {
		t.Fatalf("http request error: %s", err)
	}
	rr := httptest.NewRecorder()

	app.Server.Mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d\nbody: %s", rr.Code, rr.Body.String())
	}
	var got []dto.LocationResponse
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("faild to decode response: %v", err)
	}

	expected := domain.ErrorResponse{
		Message: "Некорректное значение в id",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %+v, got %+v", expected, got)
	}
}

func TestSearchLocation_error_locationNotFound(t *testing.T) {
	db := testutils2.NewTestDB()
	app := testutils2.NewTestApp(db)

	if err := seeds.AddUser(db); err != nil {
		t.Fatalf("data seed error %s", err)
	}

	req, err := http.NewRequest(http.MethodGet, "/searchLocation/1", nil)
	if err != nil {
		t.Fatalf("http request error: %s", err)
	}
	rr := httptest.NewRecorder()

	app.Server.Mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d\nbody: %s", rr.Code, rr.Body.String())
	}
	var got []dto.LocationResponse
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("faild to decode response: %v", err)
	}

	expected := domain.ErrorResponse{
		Message: "Данная локация не найдена",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %+v, got %+v", expected, got)
	}
}

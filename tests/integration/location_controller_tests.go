package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"weather-viewer/internal/domain"
	"weather-viewer/tests/seeds"
	"weather-viewer/testutils"
)

func TestSearchLocation_success(t *testing.T) {
	db := testutils.NewTestDB()
	app := testutils.NewTestApp(db)

	if err := seeds.AddUser(db); err != nil {
		t.Fatalf("data seed error %s", err)
	}

	if err := seeds.AddLocations(db); err != nil {
		t.Fatalf("data seed error %s", err)
	}
	req, err := http.NewRequest(http.MethodGet, "/searchLocation", nil)
	if err != nil {
		t.Fatalf("http request error: %s", err)
	}
	rr := httptest.NewRecorder()

	app.Server.Mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d\nbody: %s", rr.Code, rr.Body.String())
	}
	var got []domain.LocationResponse
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("faild to decode response: %v", err)
	}
	expected := []domain.LocationResponse{
		{Name: "Москва", UserID: 1, Latitude: 0, Longitude: 0},
		{Name: "Санкт-Петербург", UserID: 1, Latitude: 1, Longitude: 1},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected %+v, got %+v", expected, got)
	}
}

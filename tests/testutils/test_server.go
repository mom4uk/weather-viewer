package testutils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
)

func NewErrorServer(status int) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)

		location := r.URL.Query().Get("q")

		switch location {
		case "Москва":
			mskWeather, err := os.ReadFile("../../tests/fixtures/msk_weather_fixture.json")
			if err != nil {
				panic(err)
			}
			_, err = fmt.Fprintf(
				w,
				`%s`,
				string(mskWeather),
			)
			if err != nil {
				panic(err)
			}
		case "Санкт-Петербург":
			sbpWeather, err := os.ReadFile("../../tests/fixtures/spb_weather_fixture.json")
			if err != nil {
				panic(err)
			}
			_, err = fmt.Fprintf(
				w,
				`%s`,
				string(sbpWeather),
			)
			if err != nil {
				panic(err)
			}
		default:
			_, err := fmt.Fprintf(w, `На эту локацию нет стаба %s`, location)
			if err != nil {
				panic(err)
			}
			return
		}
	}))
}

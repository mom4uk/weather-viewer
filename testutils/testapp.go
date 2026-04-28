package testutils

import (
	"weather-viewer/server"
)

type TestApp struct {
	DB     *TestDB
	Server *server.Server
}

func NewTestApp(db *TestDB) *TestApp {
	return &TestApp{
		DB:     db,
		Server: server.NewServer(),
	}
}

package seeds

import (
	"weather-viewer/internal/testutils"
)

func AddUser(testDb *testutils.TestDB) error {
	_, err := testDb.DB.Exec(`
		INSERT INTO users (login, password) VALUES ('test', 'test')
	`)
	if err != nil {
		return err
	}
	return nil
}

func AddLocations(testDb *testutils.TestDB) error {
	_, err := testDb.DB.Exec(`
		INSERT INTO locations (name, user_id, latitude, longitude) VALUES ('Москва', 1, 0, 0);
		INSERT INTO locations (name, user_id, latitude, longitude) VALUES ('Санкт-Петербург', 1, 1, 1);
	`)
	if err != nil {
		return err
	}
	return nil
}

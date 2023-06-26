package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/AhmedEnnaime/GoTest/config"
)

var a config.App

func TestMain(m *testing.M) {

	a.Initialize(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()

	code := m.Run()
	clearTable()
	os.Exit(code)

}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr

}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
    id SERIAL,
    name TEXT NOT NULL,
    age INTEGER,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)`

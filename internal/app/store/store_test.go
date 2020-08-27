package store

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
)

// Test mock-DB
func TestInsertDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Init store and log instance
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)

	st := NewStore(logger, db)
	info := AppInfo{1, "Stellaris"}

	// Init expected SQL query to mock DB
	mock.ExpectExec("INSERT INTO apps").
		WithArgs(info.AppID, info.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Now we execute our method
	err = st.InsertAppInfo(&info)
	if err != nil {
		t.Errorf("Error was not expected while updating stats: %s", err)
	}

	// We make sure that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

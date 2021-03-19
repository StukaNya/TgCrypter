package store

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	uuid "github.com/satori/go.uuid"
)

// Test mock-DB
func TestInsertDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Init store and context
	ctx := context.Background()

	st := NewPinCodeStorage(db)
	err = st.InitTable(ctx)
	if err != nil {
		t.Fatalf("unable to init sql mock table: %v", err)
	}
	info := PinMock{
		id:   uuid.NewV4(),
		data: []byte{'0', 'x', 'd', 'e', 'a', 'd'},
	}

	// Init expected SQL query to mock DB
	mock.ExpectExec("INSERT INTO pin_code").
		WithArgs(info.id, info.data).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Now we execute our method
	err = st.StorePinHash(ctx, info.id, info.data)
	if err != nil {
		t.Errorf("Error was not expected while updating stats: %s", err)
	}

	// We make sure that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

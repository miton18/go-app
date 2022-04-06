package postgresql

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	err := SetupLocalDB(context.Background())
	if err != nil {
		t.Fatalf(err.Error())
	}

	db, err := New(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if err := db.PingContext(context.Background()); err != nil {
		t.Fatal(err)
	}
}

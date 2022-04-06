package postgresql

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func New(ctx context.Context) (*sql.DB, error) {
	uri := GetURI()
	if uri == "" {
		return nil, errors.Errorf("your environment does not provide valide PostgreSQL connection settings")
	}

	db, err := otelsql.Open(
		"postgres",
		uri,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithDBName(GetDatabase()),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, errors.Wrap(err, "database connection tests failed")
	}

	return db, nil
}

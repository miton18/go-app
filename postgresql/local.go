package postgresql

import (
	"context"
	"fmt"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
)

func SetupLocalDB(ctx context.Context) error {
	user := "test"
	password := "test"
	database := "test"

	req := testcontainers.ContainerRequest{
		Image:        "postgres:14.2-bullseye",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_DB":       database,
			"POSTGRES_USER":     user,
			"POSTGRES_PASSWORD": password,
		},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return err
	}

	host, err := container.Host(ctx)
	if err != nil {
		return err
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return err
	}

	os.Setenv(POSTGRESQL_ADDON_URI, fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port.Int(), database))

	return nil
}

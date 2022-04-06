package postgresql

import "os"

const (
	POSTGRESQL_ADDON_DB       = "POSTGRESQL_ADDON_DB"
	POSTGRESQL_ADDON_PASSWORD = "POSTGRESQL_ADDON_PASSWORD"
	POSTGRESQL_ADDON_ROLE     = "POSTGRESQL_ADDON_ROLE"
	POSTGRESQL_ADDON_USER     = "POSTGRESQL_ADDON_USER"
	POSTGRESQL_ADDON_URI      = "POSTGRESQL_ADDON_URI"
)

// Get Database URI
func GetURI() string {
	return os.Getenv(POSTGRESQL_ADDON_URI)
}

func GetUser() string {
	return os.Getenv(POSTGRESQL_ADDON_USER)
}

func GetPassword() string {
	return os.Getenv(POSTGRESQL_ADDON_PASSWORD)
}

func GetDatabase() string {
	return os.Getenv(POSTGRESQL_ADDON_DB)
}

func GetRole() string {
	return os.Getenv(POSTGRESQL_ADDON_ROLE)
}

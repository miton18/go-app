package app

import (
	"os"
	"strings"

	"github.com/gofrs/uuid"
)

// Internal id of current deployment
func GetDeploymentID() uuid.UUID {
	return uuid.FromStringOrNil(os.Getenv(DEPLOYMENT_ID))
}

// The id of the commit that’s currently running
func GetCommitID() string {
	return os.Getenv(COMMIT_ID)
}

func IsDebugEnabled() bool {
	ts := strings.ToLower(os.Getenv(TROUBLESHOOT))

	if ts == "true" || ts == "yes" {
		return true
	}

	return false
}

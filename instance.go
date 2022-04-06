package app

import (
	"os"
	"strconv"

	"github.com/gofrs/uuid"
)

// Allows your application to differentiate each running node on the application level.
// 0, 1…
func GetInstanceNumber() int {
	str := os.Getenv(INSTANCE_NUMBER)
	if str == "" {
		return 0
	}

	n, _ := strconv.Atoi(str)
	return n
}

// Whether this instance is a “build” instance or a “production” instance.
func GetInstanceType() string {
	return os.Getenv(INSTANCE_TYPE)
}

// The ID of the current instance (scaler) of your application. It’s unique for each instance of your application and changes every time you deploy it.
func GetInstanceID() uuid.UUID {
	return uuid.FromStringOrNil(os.Getenv(INSTANCE_ID))
}

// The random generated string as instance pretty name using pokemon names.
func GetInstancePrettyName() string {
	return os.Getenv(PRETTY_INSTANCE_NAME)
}

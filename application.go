package app

import "os"

// The ID of your Clever Cloud application
func GetAppID() string {
	return os.Getenv(APP_ID)
}

// The customer defined application name
func GetAppName() string {
	return os.Getenv(APP_NAME)
}

// The absolute path to your application folder
func GetAppHome() string {
	return os.Getenv(APP_HOME)
}

// Folder in which the application is located (inside the git repository)
func GetAppDirectory() string {
	return os.Getenv(APP_FOLDER)
}

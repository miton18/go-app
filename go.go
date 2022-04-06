package app

import "os"

func GetGoPKG() string {
	return os.Getenv(GO_PKG)
}

func GetGoBuildTool() string {
	return os.Getenv(GO_BUILD_TOOL)
}

func GetGoRunDirectory() string {
	return os.Getenv(GO_RUNDIR)
}

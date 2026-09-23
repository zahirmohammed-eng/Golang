package main


// 2. Hello cloud native
// vgdsgh
import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	appName := "Cloud-Native Service"
	appVersion := "v1.0.0"
	goVersion := runtime.Version()
	envName := os.Getenv("APP_ENV")

	if envName == "" {
		envName = "development" // default fallback
	}

	fmt.Printf("Application Name:    %s\n", appName)
	fmt.Printf("Application Version: %s\n", appVersion)
	fmt.Printf("Go Version:          %s\n", goVersion)
	fmt.Printf("Environment Name:    %s\n", envName)
}

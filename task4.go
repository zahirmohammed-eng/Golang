
// 4. Go Toolchain Practice

package main

go run main.go       // Compiles and runs Go source files directly in memory
go build             // Compiles package dependencies and source files into an executable binary
go fmt ./...         // Automatically formats Go code according to standard style guidelines
go vet ./...         // Examines Go source code and reports suspicious constructs/bugs
go test ./...        // Runs unit tests found in *_test.go files within the package

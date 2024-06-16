//usage: go generate makefile.go

//go:build ignore

package makefile

//go:generate gofmt -e -d -s -w .
//go:generate go mod tidy -v
//go:generate go build -v -trimpath -o bin/ ./...
//go:generate go test -v -vet all ./...

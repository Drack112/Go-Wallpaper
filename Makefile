build_windows:
	GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe main.go

build_osx:
	GOOS=darwin GOARCH=amd64 go build -o bin/app-amd64-darwin main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux main.go

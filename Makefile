build_windows:
	GOOS=windows GOARCH=amd64 go build -o bin/wallpaper-flare-go-exe main.go

build_osx:
	GOOS=darwin GOARCH=amd64 go build -o bin/wallpaper-flare-go-darwin main.go

build_linux:
	GOOS=linux GOARCH=amd64 go build -o bin/wallpaper-flare-go-linux main.go

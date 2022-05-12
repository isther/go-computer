PROJECT = computer
VERSION = v0.0.1
MAIN_PATH="main.go"

build-linux: 
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o release/${PROJECT}-${VERSION}-linux-amd64 ${MAIN_PATH}
    
build-mac:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o release/${PROJECT}-${VERSION}-darwin-amd64 ${MAIN_PATH}

build-win:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/${PROJECT}-${VERSION}-winx64.exe ${MAIN_PATH}

build: build-linux build-mac build-win

clean: 
	@rm -rf release
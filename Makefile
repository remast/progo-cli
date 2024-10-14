.DEFAULT_GOAL=build

build:
 echo "Building cli"
 go build -o build/progo-cli .

build-windows:
 GOARCH=amd64 GOOS=windows go build -o build/progo-cli_windows .

test:
 go test ./...

run: build
 echo "Running cli"
 ./build/progo-cli print

clean:
 echo "Cleaning cli"
 go clean
 rm -rf build

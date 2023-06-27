# Bunch of commands for the cross platform compiling

env GOOS=windows GOARCH=amd64 go build -tags enterprise -o cli.exe
env GOOS=linux GOARCH=amd64 go build -tags enterprise -o cli-linux-amd64
env GOOS=windows GOARCH=amd64 go build -tags enterprise -o cli-windows-amd64.exe
env GOOS=darwin GOARCH=amd64 go build -tags enterprise -o cli-darwin-amd64
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags enterprise -o cli-linux-amd64

goreleaser - releasing with the brew


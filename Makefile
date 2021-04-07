build:
	go build -ldflags "-s" -o sts2credentials main.go

build-all:
	GOOS=windows GOARCH=amd64 go build -ldflags "-s" -o sts2credentials-windows-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s" -o sts2credentials-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -ldflags "-s" -o sts2credentials-linux-arm64 main.go
	GOOS=linux GOARCH=ppc64le go build -ldflags "-s" -o sts2credentials-linux-ppc64le main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s" -o sts2credentials-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags "-s" -o sts2credentials-darwin-arm64 main.go

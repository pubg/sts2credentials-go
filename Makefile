build:
	go build -ldflags "-s" -o sts2credentials main.go

build-all:
	goreleaser --snapshot --rm-dist

release:
	goreleaser release --rm-dist


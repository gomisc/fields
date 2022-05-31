VER = $(shell git describe --tags | sed 's/.*v\([0-9]\{1\}.[0-9]\{1,2\}.[0-9]\{1,2\}\).*/\1/')

ver:
	@echo ${VER}

test:
	go test ./...

gomod:
	go mod tidy -compat=1.16 && go mod tidy -compat=1.17
	go mod download
fmt:
	go fmt ./...

test:
	go test -v -race

bench:
	go test -bench=.

cover:
	go test -v -race -coverprofile=coverage.out

html:
	go tool cover -html=coverage.out

lint:
	golint


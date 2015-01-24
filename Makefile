fmt:
	go fmt ./...

test:
	godep go test -v -race

bench:
	godep go test -bench=.

cover:
	godep go test -v -race -coverprofile=coverage.out

html:
	go tool cover -html=coverage.out

lint:
	golint


.PHONY: tfe-browser test fmt arm64 x86_64

tfe-browser: test arm64 x86_64

test: fmt
	go test ./...

fmt:
	go fmt ./...

arm64: dist/arm64/
	GOOS="darwin" GOARCH="arm64" go build -o dist/arm64/tfe-browser ./cmd/ 

dist/arm64/: dist/
	mkdir -p dist/arm64/

x86_64: dist/x86_64/
	GOOS="darwin" GOARCH="amd64" go build -o dist/x86_64/tfe-browser ./cmd/ 

dist/x86_64/: dist/
	mkdir -p dist/

dist/:
	mkdir -p dist/
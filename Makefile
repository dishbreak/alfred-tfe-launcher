.PHONY: tfe-browser test fmt arm64 x86_64 release

version = 0.3.1

release: tfe-browser
	zip tfe-browser-$(version).alfredworkflow info.plist icon.png dist/arm64/tfe-browser dist/x86_64/tfe-browser

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
	mkdir -p dist/x86_64/

dist/:
	mkdir -p dist/
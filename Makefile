.PHONY: tfe-browser test fmt arm64 x86_64 release verify_version

version = 0.2.0

release: tfe-browser
	zip tfe-browser-$(version).alfredworkflow info.plist icon.png dist/arm64/tfe-browser dist/x86_64/tfe-browser

tfe-browser: verify_version test arm64 x86_64

verify_version:
	grep "<!-- version tag --><string>$(version)</string>" info.plist

test: fmt
	go test ./...
	grep "<!-- version tag --><string>$(version)</string>" info.plist

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
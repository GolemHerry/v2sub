GOOS = darwin
GOARCH = amd64

v2sub:
	@echo "build v2sub"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build  -o build/v2sub

clean:
	@echo "clean v2sub"
	go clean -i && rm -rf build
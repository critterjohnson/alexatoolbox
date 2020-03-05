.DEFAULT_GOAL := build

build:
	@echo "building..."
	go build -o bin/main main.go

package:
	@echo "packaging..."
	zip -r  -j packages/main.zip bin/main

upload:
	@echo "uploading..."
	aws lambda update-function-code --function-name GoHello --zip-file fileb://packages/main.zip

test:
	go test ./... -v

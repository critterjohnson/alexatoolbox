.DEFAULT_GOAL := build

build:
	@echo "building..."
	go build -o bin/main main.go

package:
	@echo "packaging..."
	zip packages/main.zip bin/main

upload:
	@echo "uploading..."
	aws lambda update-function-code --function-name goask_test --zip-file fileb://main.zip

test:
	go test ./... -v

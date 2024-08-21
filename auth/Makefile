.PHONY: build
# build
build:
	@mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux go build -o ./bin/main ./cmd

.PHONY: build-x86
# build for x86
build-x86:
	@mkdir -p bin/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main ./cmd

.PHONY: run
# run
run:
	go build -o ./bin/main ./cmd && ./bin/main

.PHONY: run-hot
# run in hot-reload mode
run-hot:
	@air --build.cmd "go build -o ./bin/main ./cmd" --build.bin "./bin/main"
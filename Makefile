all: build
clean: rm bin/taskd bin/task-cli

build:
	go build -o bin/taskd cmd/daemon/main.go
	go build -o bin/task-cli cmd/cli/main.go

dist:
	mkdir -p dist

	go build -ldflags="-w -s" -gcflags=all="-l" -o dist/taskd cmd/daemon/main.go
	go build -ldflags="-w -s" -gcflags=all="-l" -o dist/task-cli cmd/cli/main.go

	upx -f -q --best --lzma --brute dist/taskd
	upx -f -q --best --lzma --brute dist/task-cli

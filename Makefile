.PHONY: build clean

build:
	go build -o run cmd/api/main.go

clean:
	rm -f run

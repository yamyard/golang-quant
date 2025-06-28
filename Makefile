.PHONY: build clean

build:
	go build -o golang-quant cmd/api/main.go

clean:
	rm -f golang-quant

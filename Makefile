.PHONY: build 

build:
	mkdir -p build
	go build -o build/fyi main.go

clean:
	rm -rf build
	
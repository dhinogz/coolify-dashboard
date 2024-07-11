run: build
	./bin/web serve

build: generate
	go build -o ./bin/web .

generate:
	templ generate

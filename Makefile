.PHONY:run
run: build
	@./bin/app
build:
	@go build -o bin/app .
tailwind: 
	@bun watch
templ:
	@templ generate --watch --proxy=http://localhost:8000
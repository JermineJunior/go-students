GORUN = go run
main_package_path = ./
binary_name = student_managment
build_dir = ./tmp
# goose conf
db_driver = sqlite3
db_dir = ./internal/db/migrations
db_file = app.db
AIR = air
.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
## tailwind : compile tailwind css to the input file
.PHONY: tailwind
tailwind:
	@bun dev
## watch : watch changes and update the styles
.PHONY: watch
watch:
	@bun watch
## dev: run the development server
.PHONY: dev
dev:
	${AIR}
## templ : compile the templ files and proxy the build server
.PHONY: templ
templ:
	@templ generate --watch --proxy=http://localhost:8000
## build : build the application
.PHONY: build
build:
	@mkdir -p ${build_dir}
	GOARCH=amd64 GOOS=linux go build -o ${build_dir}/${binary_name}-linux ${main_package_path}
	GOARCH=amd64 GOOS=windows go build -o ${build_dir}/${binary_name}-windows.exe ${main_package_path}
.PHONY: confirm
confirm:
	@echo -n 'Are You sure? [y/N] ' && read ans && [ $${ans:-N} = y ]
## clean : clean all previous builds
.PHONY: clean
clean: confirm
	@echo 'Cleaning up ...'
	@rm -rf ${build_dir}

## test: run tests
.PHONY: test
test:
	${GORUN} test ./...
## format : run the go formatter
.PHONY: format
format:
	${GORUN} fmt ./...

# goose targets
## migrate : migrate the database
.PHONY: migrate
migrate:
	@goose --dir${db_dir} ${db_driver} ${db_file} up
## rollback : rollback the latest migration
.PHONY: roolback
roolback:
	@goose --dir${db_dir} ${db_driver} ${db_file} down
## migration : create a migration file
.PHONY: migration
migration:
	@goose --dir${db_dir} create ${name} sql

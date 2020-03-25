APP=palu-covid
APP_EXECUTABLE="./out/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

build:
	mkdir -p ./out && go build  -o ${APP_EXECUTABLE}

run:
	${APP_EXECUTABLE} server

migrate:
	${APP_EXECUTABLE} migrate

test:
	go test -v ./...

copy-config:
	@echo "Copying sample.config.yml to config.yml ..."
	cp sample.config.yml config.yml

build-run: build run

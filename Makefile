
include .env
export

all: build-docker-update build-dev start

build-docker-update:
	docker-compose pull
	docker-compose build

build-dev:
	bin/exec dep ensure

build-prod:
	docker build -t golang-docker-bootstrap -f docker/prod/Dockerfile --build-arg PACKAGE_NAME .
	@echo "\nDone. usage : docker run --rm golang-docker-bootstrap"

start:
	docker-compose up -d --force-recreate

logs:
	docker-compose logs -f

stop:
	docker-compose stop

clean: stop
	docker-compose down -v --remove-orphans || true

restart: clean start logs

test-cs:
	@[ -z `docker-compose run --rm cli gofmt -s -l cmd internal` ] || (docker-compose run --rm cli gofmt -s -d -e cmd internal && false)
	@bin/exec golint -set_exit_status ${PACKAGE_NAME}/cmd/...
	@bin/exec golint -set_exit_status ${PACKAGE_NAME}/internal/...
	@echo "cs ok!"

fix-cs:
	bin/exec gofmt -s -l -w cmd internal
	bin/exec golint ${PACKAGE_NAME}/cmd/...
	bin/exec golint ${PACKAGE_NAME}/internal/...

test: test-cs

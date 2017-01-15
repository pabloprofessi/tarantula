export ROOT_DIR=${PWD}

build:
	go install github.com/pprofessi/server/

run:
	docker-compose up

br:
	go install github.com/pprofessi/server/
	docker-compose up tarantula

stop:
	docker-compose down

init:
	export GOPATH=$GOPATH:/home/pprofessi/projects/tarantula
	docker-compose up -d db
export ROOT_DIR=${PWD}


run:
	docker-compose up

br:
	go install github.com/pprofessi/server/
	docker-compose up tarantula

stop:
	docker-compose down

init:
	docker-compose up -d db

bash:
	docker exec -ti tarantula_tarantula_1 bash

db:
	mysql -u tarantula -h 172.17.0.2 -ptarantula tarantula

runprod:
	docker run -p 8080:8080 -e "ENV=prod" taramtula-proxy-prod 

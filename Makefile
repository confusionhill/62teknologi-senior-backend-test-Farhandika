services-up:
	@echo "SVCNAME=${SVCNAME}"
	docker-compose -f docker-compose.yaml up --build ${SVCNAME}
services-stop:
	docker compose stop
services-down:
	docker compose docker-compose.yaml down
init:
	go mod vendor
	echo "Project initiated"
rest-run:
	go run application/rest/*.go
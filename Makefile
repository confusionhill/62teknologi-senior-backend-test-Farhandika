rest-up:
	docker compose up -d
rest-stop:
	docker compose stop
rest-down:
	docker compose down
init:
	go mod vendor
	echo "Project initiated"
rest-run:
	go run application/rest/*.go
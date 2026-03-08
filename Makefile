up:
	docker compose up --build -d
down:
	docker compose down
restart:
	make down && make up
container:
	docker exec -it wallace-app sh

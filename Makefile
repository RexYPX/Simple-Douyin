.PHONY: start stop down

start:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down

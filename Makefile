.PHONY: start stop down

start:
	docker-compose up -d

stop:
	docker-compose stop

down:
	docker-compose down

run:
	cd cmd/api
	go build
	cd ../..
	sh cmd/comment/build.sh
	sh cmd/favorite/build.sh
	sh cmd/feed/build.sh
	sh cmd/message/build.sh
	sh cmd/publish/build.sh
	sh cmd/relation/build.sh
	sh cmd/user/build.sh
	./cmd/api/api
	sh cmd/comment/output/bootstrap.sh &
	sh cmd/favorite/output/bootstrap.sh &
	sh cmd/feed/output/bootstrap.sh &
	sh cmd/message/output/bootstrap.sh &
	sh cmd/publish/output/bootstrap.sh &
	sh cmd/relation/output/bootstrap.sh &
	sh cmd/user/output/bootstrap.sh &
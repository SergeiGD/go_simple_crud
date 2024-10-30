build:
	docker-compose up --build

run:
	docker-compose up

stop:
	docker-compose stop

make_migration:
	goose -dir=migrations create $(NAME) sql
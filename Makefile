localstack:
	docker-compose up -d
localstack-rebuild:
	docker-compose up -d --build
buld-image:
	docker-compose -f ./docker-compose-build.yml down -v
	docker-compose -f ./docker-compose-build.yml build --no-cache
run-image:
	docker-compose -f ./docker-compose-build.yml up
run-rebuild-image:
	docker-compose -f ./docker-compose-build.yml down -v
	docker-compose -f ./docker-compose-build.yml up --build
stop:
	docker-compose -f ./docker-compose-build.yml stop

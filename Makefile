localstack:
	docker-compose up -d
localstack-rebuild:
	docker-compose up -d --build
build-image:
	docker-compose -f ./docker-compose-build.yml down -v
	docker-compose -f ./docker-compose-build.yml build --no-cache
run-image:
	docker-compose -f ./docker-compose-build.yml up
run-rebuild-image:
	docker-compose -f ./docker-compose-build.yml down -v
	docker-compose -f ./docker-compose-build.yml up --build
stop:
	docker-compose -f ./docker-compose-build.yml stop
deploy:
	docker build -t localstack-web-ui .
	docker tag localstack-web-ui:latest aolb/localstack-web-ui:latest
	docker push aolb/localstack-web-ui:latest
start-local-dev:
	docker-compose start
	php artisan serve > /dev/null 2>&1 &
	npm --prefix ./frontend/ start > /dev/null 2>&1 &

update-laravel:
	composer update --with-all-dependencies
	php artisan migrate

update-react:
	cd frontend && npm update
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
	cd frontend && npm run build
	docker build -t localstack-web-ui .
	docker tag localstack-web-ui:latest aolb/localstack-web-ui:latest
	docker push aolb/localstack-web-ui:latest
start-local-dev:
	docker-compose start
	npm --prefix ./frontend/ start > /dev/null 2>&1 &
	cd go-api && go run .
update-react:
	cd frontend && npm update
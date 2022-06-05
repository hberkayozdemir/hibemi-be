up:
	sh .dev/up.sh

down:
	sh .dev/down.sh

build:
	go build main.go

run:
	APP_ENV=dev go run ./cmd/foo-api/main.go

lint:
	golangci-lint run -c .dev/.golangci.yml

unit-test:
	go test ./... -short

consumer-test:
	go test ./... -run TestConsumer

provider-test:
	go test -v `go list ./...|grep contract`

contract-test: consumer-test provider-test

load-test:
	docker run -i loadimpact/k6 run --include-system-env-vars --out influxdb=http://${K6_INFLUXDB_USERNAME}:${K6_INFLUXDB_PASSWORD}@${INFLUX_DB_URL} \
	-e LOAD_TEST_URL=${LOAD_TEST_URL} - <load-test.js

all-tests: unit-test contract-test load-test

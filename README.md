[![pipeline status](https://gitlab.com/modanisatech/marketplace/golang-service-template/badges/master/pipeline.svg)](https://gitlab.com/modanisatech/marketplace/golang-service-template/-/commits/master)    [![coverage report](https://gitlab.com/modanisatech/marketplace/golang-service-template/badges/master/coverage.svg)](https://gitlab.com/modanisatech/marketplace/golang-service-template/-/commits/master)

# Go Lang Service Template

## Development

Run project with all dependencies

```
make up
```

Shutdown project

```
make down
```

Build the project

```
make build
```

Run the project

```
make run
```

Setup githooks

```
git config core.hooksPath config/git/hooks/
```

### Testing

Run unit tests

```
make unit-test
```

Run consumer test

```
make consumer-test
```

Run provider test

```
make provider-test
```

Run consumer and provider tests

```
make contract-test
```

Run all tests

```
make all-tests
```

Run load tests

Before running the load tests, you need to set LOAD_TEST_URL environment variable.

```
docker run -i loadimpact/k6 run - <load-test.js
```

### Accessing Service

Service will be automatically forwarded to `localhost:3001`

### Accessing Elastic

```
kubectl port-forward --namespace default svc/dev-elastic-elasticsearch-coordinating-only 9200:9200
```

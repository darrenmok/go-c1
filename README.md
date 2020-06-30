# go-c1
For a challenge

## Development

Install GO, and GO!

## Basic requirements

1. Require docker, install [here](https://docs.docker.com/get-docker/)
1. Require docker-compose, install [here](https://docs.docker.com/compose/install/)

## Test

Need to run redis before running any test, perform the following
```bash
docker run -p 6379:6379 -d --rm --name redis redis:6-alpine
export REDIS_URL=redis://localhost:6379
go test -v ./...
```

Remember to cleanup at the end
```bash
docker kill redis
```

## Run

Just run the following commands
```bash
docker-compose up -d
```

The you should be able to access http://localhost:3000

And as always to cleanup
```bash
docker-compose down
```
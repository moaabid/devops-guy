# Introduction to Learning Go

# Run Go in Docker

We can also run go in a small docker container: <br/>

```
cd golang\introduction

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
go version

```

## Development environment

```
FROM golang:1.15 as dev

WORKDIR /work
```

## Building our code

```
FROM golang:1.15 as build

WORKDIR /app
COPY ./app/* /app/
RUN go build -o app
```

## The Runtime

```
FROM alpine as runtime 
COPY --from=build /app/app /
CMD ./app
```

## Building the Container

```
docker build --target build . -t customer-app

```
## Building the Container runtime

```
docker build . -t customer-app

```


## Running the Container

```
docker run customer-app
```



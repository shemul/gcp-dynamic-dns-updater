FROM golang:alpine as build

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk add --update curl && rm -rf /var/cache/apk/*
RUN mkdir -p /app

WORKDIR /app
COPY app /app

FROM alpine
WORKDIR /app
COPY --from=build /app/app .
ENTRYPOINT ./app
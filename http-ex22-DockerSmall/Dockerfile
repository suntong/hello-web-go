# golang:alpine as build-env
FROM golang:alpine AS build-env

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN cd /app && go build -o myapp .

FROM alpine
RUN mkdir /app
COPY --from=build-env /app/myapp /app

EXPOSE 8080
ENTRYPOINT ["/app/myapp"]

# golang:latest as build-env
FROM golang:latest AS build-env

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN cd /app && GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp .
# go build -o myapp

FROM scratch
COPY --from=build-env /app/myapp /app/images/* /

EXPOSE 8080
ENTRYPOINT /myapp

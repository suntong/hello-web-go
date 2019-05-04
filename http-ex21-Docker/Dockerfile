# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest
# create a working directory
WORKDIR /go/src/app
# add source code
ADD src src
# run main.go
CMD ["go", "run", "src/main.go"]
# http server listens on port 8080.
#EXPOSE 8080

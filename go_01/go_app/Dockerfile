FROM golang:latest
LABEL maintainer "linnefromice"
COPY . /go/src/app
WORKDIR /go/src/app
RUN go mod download
WORKDIR /go/src/app/src
EXPOSE 5000
CMD [ "go", "run", "main.go" ]
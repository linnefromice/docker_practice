FROM golang:latest
LABEL maintainer "linnefromice"
COPY . /go/src/app
WORKDIR /go/src/app
RUN go build .
EXPOSE 5000
CMD [ "./go_server" ]
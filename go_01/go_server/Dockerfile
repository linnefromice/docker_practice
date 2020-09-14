FROM golang:latest
LABEL maintainer "linnefromice"
COPY . /go/src/app
WORKDIR /go/src/app
RUN chmod 755 ./wait-for-it.sh
RUN go build .
EXPOSE 5000
CMD [ "./go_server" ]
FROM golang:latest

COPY /Api/ /go/src/api

WORKDIR /go/src/api/

RUN go mod download
RUN apt-get update
RUN apt-get install -y  mariadb-client
#&& apt-get install -y default-mysql-client

#CMD ["go", "run", "main.go"]
CMD ["sh", "./start.sh"]
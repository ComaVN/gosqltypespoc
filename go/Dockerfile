FROM golang:1.11.0

RUN go get -d \
    github.com/go-sql-driver/mysql \
;
COPY *.go ./
RUN go build -o poc
CMD ["./poc"]

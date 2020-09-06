FROM golang:alpine
WORKDIR /go/src/go-dbash
COPY . /go/src/go-dbash
RUN go build -o /go/usr/go-dbash .
EXPOSE 8000
CMD ["/go/usr/go-dbash"]

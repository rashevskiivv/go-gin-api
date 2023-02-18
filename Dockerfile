FROM golang
WORKDIR /go/src/go-gin-api
COPY . .
RUN go mod download
COPY . /go/src/go-gin-api
RUN go build -o bin/server ./cmd/main.go
EXPOSE ${PORT} ${PORT}
CMD ["./bin/server"]
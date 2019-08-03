FROM golang:latest


ADD . /app/
WORKDIR /app

RUN go build main.go

ENTRYPOINT [ "go", "run", "main.go" ]

FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go get github.com/cosmtrek/air@v1.29.0

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app","air","main.go"]
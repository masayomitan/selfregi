FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify
RUN go get -d github.com/gorilla/mux
RUN go get -u github.com/cosmtrek/air@v1.29.0 && \
    go build -o /go/bin/air github.com/cosmtrek/air
RUN go get -d github.com/gorilla/mux
RUN go get github.com/joho/godotenv

CMD ["air"]
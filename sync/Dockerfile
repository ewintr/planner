FROM golang:1.23-alpine
RUN apk add --no-cache git

WORKDIR /src
COPY . ./

RUN go mod download

RUN go build -o syncservice

CMD ["./syncservice"]

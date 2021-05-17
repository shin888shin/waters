# FROM golang:1.16 AS builder
# WORKDIR /go/src/github.com/shin888shin/waters
# RUN mkdir /go/src/github.com/shin888shin/waters
# ADD . /go/src/github.com/shin888shin/waters

# RUN go mod download
# RUN go build -o main .
# CMD ["/go/src/github.com/shin888shin/waters/main"]


FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /usr/app
COPY --from=builder /app/main .

EXPOSE 8000

# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
# RUN chmod +x /wait

CMD ["./main"]

# docker build -t shin888shin/waters:latest .

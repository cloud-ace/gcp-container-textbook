FROM golang:latest as builder

WORKDIR /app

COPY go.mod main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates curl

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 80

CMD ["./main"]

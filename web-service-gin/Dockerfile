FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o student-api .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/student-api .

EXPOSE 8080

CMD ["./student-api"]
FROM golang as builder

WORKDIR /app

COPY go.mod ./

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o assign-aws .

FROM alpine:latest

COPY --from=builder /app/assign-aws /assign-aws

CMD ["./assign-aws"]

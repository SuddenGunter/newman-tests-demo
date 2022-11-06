FROM golang:1.19.0-alpine3.16 as builder

ENV CGO_ENABLED 0

WORKDIR /src

COPY . ./

RUN go build -o /service

FROM alpine:3.16

WORKDIR /
COPY --from=builder /service .

CMD ["/service"]
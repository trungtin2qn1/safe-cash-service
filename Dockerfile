FROM golang:1.14-alpine as builder
WORKDIR /go/src/safe-cash-service
COPY . .
RUN apk add --update git make
RUN echo $GO111MODULE
RUN go mod vendor
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/safe-cash-service/safe-cash-service ./safe-cash-service
COPY --from=builder /go/src/safe-cash-service/keys ./keys
CMD ["/safe-cash-service"]
EXPOSE 5000
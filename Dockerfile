FROM golang:1.12-alpine as builder
WORKDIR /go/src/safe-cash-service
COPY . .
RUN apk add --update git make
RUN go get -u github.com/Masterminds/glide
RUN glide install
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/safe-cash-service/safe-cash-service ./safe-cash-service
CMD ["/safe-cash-service"]
EXPOSE 6000
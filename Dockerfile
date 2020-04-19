FROM golang:1.14-alpine as builder
WORKDIR /go/src/safe-cash-service
COPY . .
RUN apk add --update git make
# RUN go get -u github.com/gin-gonic/gin
# RUN go get -u github.com/Masterminds/glide
# RUN glide install
RUN go mod vendor
# RUN go get ./...
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/safe-cash-service/safe-cash-service ./safe-cash-service
COPY --from=builder /go/src/safe-cash-service/keys ./keys
CMD ["/safe-cash-service"]
EXPOSE 6000
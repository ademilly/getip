FROM golang as builder
WORKDIR /go/src/github.com/ademilly/getip
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine
WORKDIR /root/
COPY --from=builder /go/src/github.com/ademilly/getip/getip .
CMD ["./getip"]

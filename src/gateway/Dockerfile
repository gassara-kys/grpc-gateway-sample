FROM golang AS builder
WORKDIR /app
ADD *.go go.* ./
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o gateway .

FROM alpine
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/gateway /usr/local/bin/
ENV PORT= \
    GREETING_SVC_ADDR= \
    TZ=Asia/Tokyo
WORKDIR /usr/local
CMD ["bin/gateway"]

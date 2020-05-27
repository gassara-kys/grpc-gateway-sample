FROM golang AS builder
WORKDIR /app
ADD *.go go.* ./
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o greeting .

FROM alpine
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /app/greeting /usr/local/bin/
ENV PORT= \
    TZ=Asia/Tokyo
WORKDIR /usr/local
CMD ["bin/greeting"]

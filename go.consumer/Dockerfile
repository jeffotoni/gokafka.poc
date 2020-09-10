FROM golang:1.15 as builder
WORKDIR /go/src/go.consumer
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o go.consumer main.go
RUN cp go.consumer /go/bin/go.consumer
RUN cp config.yaml /go/bin/config.yaml
RUN ls -lh

FROM alpine:latest AS final
RUN apk update
RUN apk add --no-cache tzdata
ENV TZ="America/Sao_Paulo"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/bin/go.consumer /
COPY --from=builder /go/bin/config.yaml /

CMD ["/go.consumer"]
FROM golang:1.15 as builder
WORKDIR /go/src/server.game.blocks
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o server.game.blocks main.go
RUN cp server.game.blocks /go/bin/server.game.blocks
RUN cp index.html /go/bin/index.html
RUN cp wasm_exec.js /go/bin/wasm_exec.js
RUN cp webassembly.wasm /go/bin/webassembly.wasm
RUN ls -lh

FROM alpine:latest AS final
RUN apk update
RUN apk add --no-cache tzdata
ENV TZ="America/Sao_Paulo"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/bin/server.game.blocks /
COPY --from=builder /go/bin/index.html /
COPY --from=builder /go/bin/wasm_exec.js /
COPY --from=builder /go/bin/webassembly.wasm /

CMD ["/server.game.blocks"]
# one
FROM golang:1.17 as one
WORKDIR /src/
ENV GOPROXY=https://proxy.golang.com.cn,direct
ENV GIN_MODE=release
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN make build
# two
FROM debian:10-slim
ENV PORT=8080
WORKDIR /
COPY --from=one /src/dist/dragonfly ./
EXPOSE $PORT
ENTRYPOINT ["/dragonfly", "serve"]

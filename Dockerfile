# build stage
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git make gcc g++ ca-certificates && update-ca-certificates

WORKDIR /src

COPY . .

RUN make tidy

RUN make bin/now_linux_amd64

# final stage
FROM gcr.io/distroless/base:latest
COPY --from=builder /src/bin/now_linux_amd64 /bin/now
ENTRYPOINT ["/bin/now"]
CMD ["--help"]

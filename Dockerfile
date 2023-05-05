FROM golang:1.20.3-alpine3.17 AS builder

ENV SERVICE=vk-chatbot

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/$SERVICE

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git
RUN apk add protobuf-dev

# Build Go binary
COPY . /go/src/$SERVICE/
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,mode=0755,target=/go/pkg/mod make tidy build

# Deployment container
FROM scratch

ENV SERVICE=vk-chatbot

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/$SERVICE/$SERVICE /$SERVICE
EXPOSE 8080
ENTRYPOINT ["/vk-chatbot"]
CMD []

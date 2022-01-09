# builder image
FROM golang:1.16 as builder
RUN mkdir /build
COPY . /build
WORKDIR /build
RUN cd /build && CGO_ENABLED=0 GOOS=linux go build -a -o server ./cmd/server/main.go


# generate clean, final image for end users
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /build .

# executable
ENTRYPOINT [ "./server" ]
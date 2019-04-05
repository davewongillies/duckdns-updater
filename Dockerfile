FROM golang:alpine as builder
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates
RUN mkdir /app
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/duckdns-updater
CMD ["/app/duckdns-updater"]

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/duckdns-updater /app/duckdns-updater
COPY --from=builder /etc/passwd /etc/passwd
USER nobody
ENTRYPOINT ["/app/duckdns-updater"]

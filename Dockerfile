FROM golang:alpine as builder
RUN mkdir /app
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /app/duckdns-updater
CMD ["/app/duckdns-updater"]

FROM scratch
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder /app/duckdns-updater /app/duckdns-updater
ENTRYPOINT ["/app/duckdns-updater"]

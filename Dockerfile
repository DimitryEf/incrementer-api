FROM golang AS builder
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o incrementer

FROM scratch
COPY --from=builder /app/incrementer /app/incrementer
EXPOSE 8080 8081
ENTRYPOINT ["/app/incrementer"]

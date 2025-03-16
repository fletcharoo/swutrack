FROM golang:latest AS build
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build ./services/swutrack/main.go

FROM debian:latest
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]
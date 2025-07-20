FROM golang:latest AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY ./service ./service
RUN go build -o /app/main ./service

FROM debian:latest
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]
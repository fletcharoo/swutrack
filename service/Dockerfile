FROM golang:latest AS build
WORKDIR /app
COPY ./service ./
RUN go mod download
RUN go build ./main.go

FROM debian:latest
COPY --from=build /app/main /app/main
ENTRYPOINT ["/app/main"]
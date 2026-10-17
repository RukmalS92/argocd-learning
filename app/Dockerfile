# Build stage
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 go build -o /go-app .

# Runtime stage: tiny image with just the binary
FROM gcr.io/distroless/static
COPY --from=build /go-app /go-app
EXPOSE 8080
ENTRYPOINT ["/go-app"]

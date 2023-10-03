# Build stage
FROM golang:1.17-alpine AS build-stage
WORKDIR /app
COPY . .
RUN go build -o main
RUN rm -rf some_unnecessary_directory

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build-stage /app/main .
EXPOSE 8080
CMD ["./main"]

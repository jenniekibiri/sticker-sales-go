# build stage
FROM golang:alpine 


WORKDIR /app

COPY . .

RUN go build -o main

RUN rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

# final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/main .

EXPOSE 8080


CMD ["./main"]

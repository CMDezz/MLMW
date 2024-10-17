# Build stage
FROM golang:1.22.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# RUN apk add curl
RUN apk add --no-cache curl \
    && apk add --no-cache bash

# Final stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY wait-for.sh .
COPY start.sh .
RUN chmod +x wait-for.sh
RUN chmod +x start.sh
COPY dto/migration ./dto/migration
# Expose port 8080 (if not already done in Dockerfile)
EXPOSE 8080

# Set the entrypoint and command
CMD ["/app/main"]
ENTRYPOINT ["/app/wait-for.sh", "postgres:5432", "--", "/app/start.sh"]
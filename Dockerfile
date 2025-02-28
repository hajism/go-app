FROM golang:1.23.6-alpine

# Update and install git
RUN apk update && apk add --no-cache git

# root directory
WORKDIR /app

# Copy File to Container
COPY . .

# Setup Timezone
RUN apk add --no-cache tzdata
ENV TZ=Asia/Jakarta

# Go Build
RUN go mod tidy
RUN go build -o main

# Expose port
EXPOSE 5000

ENTRYPOINT ["/app/main"]
######################### Stage O - Build ###############################
FROM golang:alpine3.11

# Set ENV
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Building folder
WORKDIR /build

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY="https://goproxy.io,direct"

# Copy the code into the container
COPY . .

# go mod
RUN go mod tidy
RUN go mod download

# Build
RUN go build -o main cmd/docker/main.go

# Switch to release folder
WORKDIR /dist

# Copy binary
RUN cp /build/main .

# Remove /build folder 
RUN rm -rf /build

######################### Stage 1 - Production ###############################
FROM alpine:3.11.6

WORKDIR /app/

COPY --from=0 /dist/main .

# Export
EXPOSE 3000

# Start
CMD ["/app/main"]
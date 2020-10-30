FROM golang:alpine AS builder

# Move to working directory /build
WORKDIR /build

COPY go.mod .

ENV GO111MODULE=on
RUN go mod tidy

COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api .

FROM scratch

# Copy our static executable
COPY --from=builder /build/api /

#api server port
EXPOSE 8088

# Run the api binary.
ENTRYPOINT ["/api"]
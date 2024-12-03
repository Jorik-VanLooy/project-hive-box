FROM golang:1.23@sha256:ee5f0ad829b8a88be7689e04dc44eeb932857ba3299b5bb576ee2c0bab8963ff

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY /cmd/app/go.mod /cmd/app/go.sum ./
RUN go mod download

# export version as environment variable
ENV VERSION=v0.4.1

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY /cmd/app/*.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /server

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 3333

# Run
CMD ["/server"]

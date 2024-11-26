FROM golang:1.23@sha256:73f06be4578c9987ce560087e2e2ea6485fb605e3910542cadd8fa09fc5f3e31

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY /cmd/app/go.mod /cmd/app/go.sum ./
RUN go mod download

# export version as environment variable
ENV VERSION=v0.3.2

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

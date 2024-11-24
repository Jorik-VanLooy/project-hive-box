FROM golang:1.22@sha256:147f428a24c6b80b8afbdaec7f245b9e7ac342601e3aeaffb321a103b7c6b3f4

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY /cmd/app/go.mod /cmd/app/go.sum ./
RUN go mod download

# export version as environment variable
ENV VERSION=v0.1.2

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

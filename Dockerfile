# Centarus Dockerfile

# Base image
FROM ubuntu:22.04

# Download and install dependencies
RUN apt-get update && apt-get upgrade -y && apt-get install -y \
    wget

# Download and extract Go 1.19.4
RUN wget https://go.dev/dl/go1.19.4.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz

# Set environment variables
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

# Copy source code
COPY . $GOPATH/src/github.com/centarus
WORKDIR $GOPATH/src/github.com/centarus

# Build and Run
RUN go build -o centarus
CMD ["./centarus"]

# Expose port 8080
EXPOSE 8080

# End of Dockerfile
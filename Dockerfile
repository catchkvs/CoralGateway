
# Start from the latest golang base image
FROM golang:1.15.2-alpine3.12

# Add Maintainer Info
LABEL maintainer="github.com/catchkvs"
RUN apk add --no-cache --virtual .build-deps \
        curl \
        g++ \
        gcc \
        bash \
      cmake \
      sudo \
		libssh2 libssh2-dev\
		git

RUN mkdir -p /app/pkg
RUN mkdir -p /app/resources
RUN mkdir -p /app/static
COPY ./pkg /app/pkg/
COPY ./resources /app/resources
COPY ./static /app/static/
RUN ls /app/pkg/
RUN ls /app/resources

WORKDIR /app/
#ENV GOPATH=$GOPATH:/go/
RUN echo $GOPATH
# Copy go mod and sum files
COPY go.mod  /app/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
#RUN go get .
# Build the Go app
RUN go build -o server pkg/coralgateway.go

# Expose port 3040 to the outside world
EXPOSE 3030

# Command to run the executable
CMD ["./server"]
FROM golang:1.16-alpine

ENV HTTP_PORT=5000

RUN apk add --no-cache git

WORKDIR /go/src/treatfield-api

COPY ./ ./

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build the Go app
RUN go build -o /main ./

RUN ["chmod", "+x", "/go/src/treatfield-api/docker/after-build.sh"]

# Expose port 5000 to the outside world
EXPOSE 5000

ENTRYPOINT ["sh", "/go/src/treatfield-api/docker/after-build.sh"]


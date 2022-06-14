FROM golang:1.18-alpine AS builder

# Set maintainer
LABEL Maintainer="Cuong Dang Trung <cuongisstudying@gmail.com>"

# Set HEADER AND ENV FILES
ARG HEADER_FILE
ARG ENV_FILE
ENV HEADER_FILE=header_production.go
ENV ENV_FILE=.env.pro

# Check HEADER_FILE & ENV_FILE
RUN echo "File swagger: $HEADER_FILE"
RUN echo "File env: $ENV_FILE"

RUN apk add bash ca-certificates git gcc g++ libc-dev

# Set WORKDIR
RUN mkdir -p /work/spser
WORKDIR /work/spser

# Copy go.mod and go.sum
COPY go.mod .
COPY go.sum .
# RUN ls -la /work/linkepee/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# COPY everything else
COPY . /work/spser/
# RUN ls -la /work/linkepee/

# COPY $ENV_FILE /work/linkepee/.env

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go mod tidy
RUN swag init --parseDependency -g $HEADER_FILE



# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o spser .

# ######## Start a new stage from scratcsh #######
# FROM ubuntu:20.04 

# RUN apt-get update
# RUN apt-get install -y ca-certificates

# # Here we copy the rest of the source code
# WORKDIR /usr/local/bin

# COPY --from=linkepee-builder:builder /work/linkepee/. /usr/local/bin
# # COPY . /usr/local/bin

# RUN ls -la /usr/local/bin/infrastructure

CMD ["./spser"]

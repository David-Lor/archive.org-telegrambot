FROM golang:1.17.7

WORKDIR /app
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download


COPY ./src/* ./
#RUN ls -lah
RUN go build -o /tmp/built

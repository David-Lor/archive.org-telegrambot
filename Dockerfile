FROM golang:1.16 as build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY ./src/go.mod .
COPY ./src/go.sum .
RUN go mod download


COPY ./src/ ./
RUN go build -o /tmp/built


FROM scratch

COPY --from=build /tmp/built /archiveorg-telegrambot
CMD ["/archiveorg-telegrambot"]

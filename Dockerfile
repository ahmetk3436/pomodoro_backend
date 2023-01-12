FROM golang:1.16.15 AS build-env

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o main .

FROM alpine
COPY --from=build-env /app/main /usr/local/bin/
CMD ["main"]

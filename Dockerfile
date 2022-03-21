FROM golang:latest
RUN mkdir /app
WORKDIR /app
RUN mkdir ./src
RUN mkdir ./public
COPY ./public /app/public
WORKDIR /app/src
COPY ./src /app/src
RUN go mod download
RUN go build -o main ./cmd
CMD ["/app/src/main"]
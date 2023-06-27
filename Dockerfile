FROM golang:1.18
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o dumper
EXPOSE 9999
ENTRYPOINT ["/app/dumper"]
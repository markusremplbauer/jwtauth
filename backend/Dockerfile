FROM golang:1.16.3-alpine
WORKDIR /app
COPY api/go.mod api/go.sum ./
RUN go mod download
COPY api .
RUN go build -o api .
EXPOSE  8080
CMD ["./api"]
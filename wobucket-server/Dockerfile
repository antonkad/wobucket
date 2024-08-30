FROM golang:1.23-alpine

RUN apk add --no-cache curl git && \
    curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
CMD ["air"]

EXPOSE 8080

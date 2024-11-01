FROM golang:1.23.2-bookworm

WORKDIR /app/webapp

RUN go install github.com/air-verse/air@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["air", "-c", ".air.toml"]
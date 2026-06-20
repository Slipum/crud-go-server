FROM golang:1.24.5-alpine

WORKDIR /app

COPY go.mod ./
#если появятся библиотеки, то добавить go.sum
RUN go mod download

COPY . .

RUN go build -o app .

EXPOSE 8080

CMD ["./app"]
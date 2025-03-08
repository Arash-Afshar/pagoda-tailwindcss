FROM golang:latest

RUN apt-get update
RUN apt-get install -y \
    nodejs \
    npm \
    sqlite3

RUN mkdir -p /app
WORKDIR /app

COPY . .

RUN make build

EXPOSE 8000

CMD ["./pagoda"]

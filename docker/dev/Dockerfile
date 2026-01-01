FROM golang:1.25-alpine

WORKDIR /app

COPY . .

RUN go install github.com/air-verse/air@latest

CMD [ "air" ]

EXPOSE 5008
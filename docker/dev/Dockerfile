FROM golang:1.25-alpine

RUN apk add --no-cache git curl

RUN go install github.com/air-verse/air@1.60.0

WORKDIR /app

COPY . .


CMD [ "air" ]

EXPOSE 5008
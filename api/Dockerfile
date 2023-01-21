FROM golang:1.19

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ./

RUN go build -o /coffe_delivery_remix_api

EXPOSE 9000

CMD [ "/coffe_delivery_remix_api" ]


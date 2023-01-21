FROM golang:1.19

WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /coffe_delivery_remix_api

CMD [ "/coffe_delivery_remix_api" ]


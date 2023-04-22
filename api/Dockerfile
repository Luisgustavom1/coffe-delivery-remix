FROM golang:1.19

WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /coffee_delivery_remix_api

CMD [ "/coffee_delivery_remix_api" ]


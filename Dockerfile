FROM golang:1.18

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o capstone

CMD ["./capstone"]

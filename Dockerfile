FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /crowdfunding

EXPOSE 3000

CMD ["/crowdfunding"]

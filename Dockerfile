FROM golang:latest
WORKDIR /app
COPY . .
COPY .env .env
RUN go build -o main cmd/http/main.go
EXPOSE 8080

# Install wait-for-it script
ADD https://github.com/vishnubob/wait-for-it/raw/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

CMD ["/wait-for-it.sh", "postgres:5432", "--", "./main"]
FROM golang:1.19

WORKDIR /app

COPY . .

RUN mkdir -p bin && cd cmd && go build -o ../bin/core_service

CMD ./bin/core_service
FROM golang:latest
WORKDIR /home/go/app
COPY . /home/go/app
CMD ["go", "run", "./main.go"]
EXPOSE 9000
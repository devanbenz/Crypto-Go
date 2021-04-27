FROM golang:latest
COPY . ./app
CMD ["go", "run", "./main.go"]
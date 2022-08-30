FROM golang:1.16

# Create app directory
WORKDIR /usr/src/app

# Copy app to directory
COPY . /usr/src/app

RUN go mod tidy

CMD ["make", "test"]

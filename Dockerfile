# Note: This image is also built in CircleCI, so limit references to internal repositories.
FROM 345594567237.dkr.ecr.us-east-1.amazonaws.com/docker-hub/library/golang:1.16

# Create app directory
WORKDIR /usr/src/app

# Copy app to directory
COPY . /usr/src/app

CMD ["make", "test"]

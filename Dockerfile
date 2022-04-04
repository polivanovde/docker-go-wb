FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app
# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go ./
COPY submitter.go ./
COPY db-worker.go ./
COPY html.go ./
COPY ./cacher/cacher.go ./cacher/
COPY ./publish/model.go ./

# Build
RUN go build -o /docker-go-wb

EXPOSE 8080
# Run
CMD [ "/docker-go-wb" ]

FROM golang:1.15-alpine

# Set destination for COPY
RUN mkdir -p /go/src/github.com/NisalSP9/Did-I-read/
WORKDIR /go/src/github.com/NisalSP9/Did-I-read/

# Download Go modules
COPY go.mod .
COPY go.sum .
COPY .env .
RUN go mod download

# Copy the source code
COPY . ./

# Build
RUN CGO_ENABLED=0 go build github.com/NisalSP9/Did-I-read/
COPY . ./
RUN chmod +x Did-I-read
CMD ["./Did-I-read"]
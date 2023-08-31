# Use an official Go runtime as a parent image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Clone the GitHub repository
#RUN git clone git@github.com:gaurilab/golang-hello.git .

# Install any necessary dependencies (if applicable)
ADD . /app


# Build the Go app
RUN go mod tidy
RUN go build -o main .

# Expose port 8080
EXPOSE 9080

# Command to run the executable
CMD ["./main"]


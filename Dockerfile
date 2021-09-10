## We specify the base image we need for our
## go application
FROM golang:1.12.0-alpine3.9

#added after error 
#//#9 0.257 go: github.com/denisenkom/go-mssqldb@v0.10.0: git init --bare in /go/pkg/mod/cache/vcs/d004f62f0bf917602489d74fdea89c7a3f58cffbf618ef1e4fc751f5d4836311: exec: "git": executable file not found in $PATH
#9 0.257 go: error loading module requirements
RUN apk add git

## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]

EXPOSE 8080 8080
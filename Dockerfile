FROM golang:latest 
RUN go get gopkg.in/urfave/cli.v1
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app
RUN go build -o api . 
CMD ["/app/api"]

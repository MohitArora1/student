FROM golang:1.12

# RUN mkdir /go/src/github.com
# RUN mkdir /go/src/github.com/MohitArora1
WORKDIR /go/src/github.com/MohitArora1/student
COPY . .

RUN go get -v ./...
RUN go build .
EXPOSE 8080

CMD ["student"]
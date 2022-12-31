FROM golang:latest

##buat folder APP
RUN mkdir /todo

##set direktori utama
WORKDIR /todo

##copy seluruh file ke completedep
ADD . /todo

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]

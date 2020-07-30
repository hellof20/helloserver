FROM golang:1.13
EXPOSE 80
COPY ./bin/helloserver /usr/local/bin/
CMD ["helloserver"]

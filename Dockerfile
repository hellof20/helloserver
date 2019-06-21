FROM golang:1.12-alpine
EXPOSE 80
COPY ./bin/hello-server /usr/local/bin/
CMD ["hello-server"]

FROM alpine

WORKDIR /web/gin

COPY ./out/linux/. .

EXPOSE 8080

CMD ./gin_hello
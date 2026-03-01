FROM alpine:3.19

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

RUN mkdir -p data

COPY ./bot .

CMD ["./bot"]
#build a smaller docker image
FROM alpine:latest

RUN mkdir /app 

COPY loggerServiceApp /app

CMD [ "/app/loggerServiceApp" ]
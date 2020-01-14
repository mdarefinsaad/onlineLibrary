FROM alpine:latest

RUN apk update

RUN apk add nginx 


RUN mkdir -p /run/nginx
RUN mv /etc/nginx/nginx.conf /etc/nginx/nginx.conf.orig

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80 

CMD ["nginx"]

# docker build . -t nginx 
# docker run -- volume 

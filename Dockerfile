FROM ubuntu:latest
LABEL authors="christophe2bu"

WORKDIR /MyTimteServer
EXPOSE 3000
COPY MyTimeServer ./
CMD ["/MyTimteServer/MyTimeServer"]
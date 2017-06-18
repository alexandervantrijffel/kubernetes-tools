#busybox-ubuntu with additional essential tools like curl, nmap, nslookup

FROM ubuntu:yakkety
MAINTAINER Alexander van Trijffel

RUN apt-get update && apt-get -y upgrade 
RUN apt-get -y --fix-missing install \
            curl \
            dnsutils \
            nmap \
            less \
            nano \
            silversearcher-ag
RUN rm -rf /var/cache/apk/*

ENV TERM=xterm


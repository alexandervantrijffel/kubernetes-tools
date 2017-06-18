# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:1.8-alpine

# RUN apk add --update bash bash-completion util-linux pciutils usbutils coreutils binutils findutils grep && rm -rf /var/cache/apk/*

# fix for golang
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

# Create a directory inside the container to store all our application and then make it the working directory.

ENV GOPATH /go:$GOPATH
ENV PATH /go/bin:$PATH

COPY gowebapp /usr/bin/
RUN chmod +x /usr/bin/gowebapp

ENV PORT 80
EXPOSE 80

ENTRYPOINT ["/usr/bin/gowebapp"]

# run with shell in alpine:
# docker run -p 80:80 -ti gowebapp sh
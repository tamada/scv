FROM alpine:3.10.1

ARG version=1.0.0

LABEL maintainer="Haruaki Tamada" \
      description="Similarities and distances calculator among vectors"

RUN    adduser -D scv \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/tamada/scv/realeases/download/v${version}/scv-${version}_linux_amd64.tar.gz \
    && curl -s -L -o scv-${version}_linux_amd64.tar.gz https://www.dropbox.com/s/3rxslv4jufwgh3r/scv-1.0.0_linux_amd64.tar.gz?dl=0 \
    && tar xfz scv-${version}_linux_amd64.tar.gz        \
    && mv scv-${version} /opt                           \
    && ln -s /opt/scv-${version} /opt/scv               \
    && ln -s /opt/scv-${version}/scv /usr/local/bin/scv \
    && rm scv-${version}_linux_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/scv"

WORKDIR /home/scv
USER    scv

ENTRYPOINT [ "/opt/scv/scv" ]

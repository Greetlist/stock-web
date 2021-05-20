FROM ubuntu:xenial

RUN set -xe \
    && adduser --disabled-password --gecos '' greetlist \
    && addgroup greetlist sudo \
    && mkdir -p /home/greetlist/workspace/

USER greetlist
WORKDIR /home/greetlist/workspace
COPY --chown=greetlist:greetlist ./web_application /home/greetlist/workspace/
CMD ["/bin/bash", "-c", "./main"]

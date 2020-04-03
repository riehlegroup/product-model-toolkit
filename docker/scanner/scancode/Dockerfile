FROM ubuntu:19.04

RUN apt-get update \
    && apt-get install -y curl python-dev xz-utils zlib1g libxml2-dev libxslt1-dev bzip2 python-pip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sL https://github.com/nexB/scancode-toolkit/releases/download/v3.1.1/scancode-toolkit-3.1.1.tar.bz2 | tar -xj \
    && mv scancode* /app

WORKDIR /app

RUN ./scancode --version
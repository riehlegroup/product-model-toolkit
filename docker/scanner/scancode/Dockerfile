# SPDX-FileCopyrightText: 2022 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
#
# SPDX-License-Identifier: Apache-2.0

FROM ubuntu:19.04

RUN apt-get update \
    && apt-get install -y curl python-dev xz-utils zlib1g libxml2-dev libxslt1-dev bzip2 python-pip \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sL https://github.com/nexB/scancode-toolkit/releases/download/v21.6.7/scancode-toolkit-21.6.7_py36-linux.tar.xz | tar -xj \
    && mv scancode* /app

WORKDIR /app

RUN ./scancode --version
# SPDX-FileCopyrightText: 2021 Cristian Mogildea
#
# SPDX-License-Identifier: Apache-2.0

FROM ubuntu:20.04

RUN apt-get update
RUN DEBIAN_FRONTEND="noninteractive" apt-get install -y curl git
RUN DEBIAN_FRONTEND="noninteractive" apt-get install -y npm
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/*

RUN git clone https://github.com/mojaloop/license-scanner /tmp/license-scanner
WORKDIR /tmp/license-scanner
RUN make build default-files
RUN make set-up

# SPDX-FileCopyrightText: 2020 Friedrich-Alexander University Erlangen-Nürnberg (FAU)
#
# SPDX-License-Identifier: Apache-2.0

FROM debian:buster-slim

RUN apt-get update \
    && apt-get install -y ruby-full \
    && apt-get install -y ruby bundler cmake pkg-config git libssl-dev zlib1g-dev curl \
    && rm -rf /var/lib/apt/lists/*

RUN gem install licensee -v 9.13.0

FROM debian:stretch-slim

RUN apt-get update \
    && apt-get install -y ruby-full \
    && apt-get install -y ruby bundler cmake pkg-config git libssl-dev zlib1g-dev\
    && rm -rf /var/lib/apt/lists/*

RUN gem install licensee -v 9.13.0

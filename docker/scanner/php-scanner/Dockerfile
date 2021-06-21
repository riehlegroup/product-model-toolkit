FROM alpine:3.13
ENV TZ=CET
RUN apk update \
 && apk add --no-cache \
        bash \
        wget \
        curl \
        ca-certificates \
        patch \
        openssh-client \
        git \
        php7-cli \
        unzip \
        curl \
        php7-tokenizer \
        php7-xml \
        php7-mbstring \
        php7-common \
        php7-curl \
        php7-gd \
        php7-intl \
        php7-json \
        php7-ldap \
        php7-mysqli \
        php7-opcache \
        php7-tidy \
        php7-zip \
        php7-pear \
        php7-phar \
 && mkdir -p /root/.ssh \
 && echo "StrictHostKeyChecking no" > /root/.ssh/config \
 && curl -sS https://getcomposer.org/installer -o composer-setup.php \
 && php composer-setup.php --install-dir=/usr/local/bin --filename=composer --1 \
 && mkdir /opt/scripts \
 && mkdir /output \
 && mkdir /source
COPY ./includes/run.sh /opt/
RUN chmod -Rf 600 /root/.ssh \
 && chmod +x /opt/*.sh
ENTRYPOINT /opt/run.sh


#!/bin/bash
export NODE_TLS_REJECT_UNAUTHORIZED=0

if [ $USE_DEFAULT_REPO -eq 1 ]; then
	cd /source
	rm -rf *
	git clone https://github.com/wikimedia/mediawiki-extensions-BlueSpiceFoundation.git /source
fi
cd /source
composer update --ignore-platform-reqs

cd /opt/scripts
git clone https://github.com/hallowelt/product-model-php-plugin.git /opt/scripts/scanner
cd /opt/scripts/scanner
composer update --ignore-platform-reqs

cp config.template.json config.json
php phpScanner.php --sourcedir=/source --outputdir=/output

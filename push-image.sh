echo $GITHUB | docker login ghcr.io -u $USERNAME --password-stdin

docker build -t phpscanner . 

docker tag phpscanner docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner:1.0.0 && docker push docker.pkg.github.com/osrgroup/product-model-toolkit/php-scanner

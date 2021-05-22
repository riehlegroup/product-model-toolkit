set -e
docker build -f docker/Dockerfile.bom -t bom .
docker build -f docker/Dockerfile.cli -t cli .
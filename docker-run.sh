set -e
docker build -f docker/bom/Dockerfile -t bom .
docker build -f docker/cli/Dockerfile -t cli .
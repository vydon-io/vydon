#!/bin/sh

###
# This script is mostly for documentation purposes or if we ever have to add another chart.
# It only needs to run a single time to claim ownership and verify the publisher.
###

docker run -it --rm \
  -v ./scripts/artifacthub/vydon-artifacthub-repo.yml:/workspace/artifacthub-repo.yml \
  ghcr.io/oras-project/oras:v1.2.0 push \
  ghcr.io/vydon-io/vydon/helm/vydon:artifacthub.io \
  -u $DOCKER_USERNAME \
  -p $DOCKER_PAT \
  --config /dev/null:application/vnd.cncf.artifacthub.config.v1+yaml \
  artifacthub-repo.yml:application/vnd.cncf.artifacthub.repository-metadata.layer.v1.yaml

docker run -it --rm \
  -v ./scripts/artifacthub/vydon-api-artifacthub-repo.yml:/workspace/artifacthub-repo.yml ghcr.io/oras-project/oras:v1.2.0 push \
  ghcr.io/vydon-io/vydon/helm/api:artifacthub.io \
  -u $DOCKER_USERNAME \
  -p $DOCKER_PAT \
  --config /dev/null:application/vnd.cncf.artifacthub.config.v1+yaml \
  artifacthub-repo.yml:application/vnd.cncf.artifacthub.repository-metadata.layer.v1.yaml

docker run -it --rm \
  -v ./scripts/artifacthub/vydon-app-artifacthub-repo.yml:/workspace/artifacthub-repo.yml ghcr.io/oras-project/oras:v1.2.0 push \
  ghcr.io/vydon-io/vydon/helm/app:artifacthub.io \
  -u $DOCKER_USERNAME \
  -p $DOCKER_PAT \
  --config /dev/null:application/vnd.cncf.artifacthub.config.v1+yaml \
  artifacthub-repo.yml:application/vnd.cncf.artifacthub.repository-metadata.layer.v1.yaml

docker run -it --rm \
  -v ~/.docker:/root/.docker \
  -v ./scripts/artifacthub/vydon-worker-artifacthub-repo.yml:/workspace/artifacthub-repo.yml ghcr.io/oras-project/oras:v1.2.0 push \
  ghcr.io/vydon-io/vydon/helm/worker:artifacthub.io \
  -u $DOCKER_USERNAME \
  -p $DOCKER_PAT \
  --config /dev/null:application/vnd.cncf.artifacthub.config.v1+yaml \
  artifacthub-repo.yml:application/vnd.cncf.artifacthub.repository-metadata.layer.v1.yaml

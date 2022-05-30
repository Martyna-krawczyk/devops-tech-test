#!/bin/sh

#get app version
version=$(git describe --always --tags 2>/dev/null)
if [ -z "$version" ]; then
  echo "no version"
  exit 1
fi

#get git commit
commit=$(git rev-list -1 HEAD 2>/dev/null)
if [ -z "$commit" ]; then
  echo "no commit"
  exit 1
fi

# build and tag docker image
docker build --build-arg "version=$version" --build-arg "commit=$commit" \
  --tag "$commit" . || { echo "failed to build docker image"; exit 1; }


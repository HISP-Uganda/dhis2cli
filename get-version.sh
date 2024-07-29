#!/bin/bash
# get-version.sh
set -e

TAG=$(git describe --tags $(git rev-list --tags --max-count=1) 2>/dev/null)
if [ -z "$TAG" ]; then
  echo "1.0.0" # default version if no tags are found
else
  echo "$TAG"
fi
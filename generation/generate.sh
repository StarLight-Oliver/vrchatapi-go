#!/usr/bin/env bash
set -euo pipefail

if [ ${#} -le 0 ]
then
  echo "Usage: generate.sh <openapi.json>" >&2
  exit 1
fi

# Generate Client
rm src docs -rf

mkdir -p src
pushd src
../node_modules/\@openapitools/openapi-generator-cli/main.js generate \
-g go \
--additional-properties=packageName=vrchatapi \
--git-user-id=StarLight-Oliver \
--git-repo-id=vrchatapi-go \
-o . \
-i "${1}" \
--http-user-agent="vrchatapi-go"
#--global-property debugOperations=true
popd

exit 0

# Remove messily pasted markdown at top of every file
find src -type f -exec sed -i '/VRChat API Banner/d' {} \;
# Remove openapi version in every file
find src -type f -exec sed -i '/The version of the OpenAPI document/d' {} \;
# Remove empty doc comments
find src -type f -exec sed -i '/^\s*\/\/\/\s*$/d' {} \;

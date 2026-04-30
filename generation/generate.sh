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

# Apply patches the OpenAPI generator can't produce on its own (see
# generation/patches/* for the rationale behind each):
#   1. strip "decoder.DisallowUnknownFields()" from model UnmarshalJSON
#   2. prefix enum constants with their type name (PUBLIC -> ReleaseStatus_PUBLIC etc.)
#   3. add the missing "time" import to api_miscellaneous.go
#   4. drop duplicate *UserPersistence* handlers from api_worlds.go
# Stages results in ./out/ for review. Swap to "-override" once you trust it
# and want to write the patched files straight to the repo root.
rm -rf out
mkdir -p out
pushd migrate
go run . -src ../src -out ../out
popd

exit 0

# Remove messily pasted markdown at top of every file
find src -type f -exec sed -i '/VRChat API Banner/d' {} \;
# Remove openapi version in every file
find src -type f -exec sed -i '/The version of the OpenAPI document/d' {} \;
# Remove empty doc comments
find src -type f -exec sed -i '/^\s*\/\/\/\s*$/d' {} \;

#!/usr/bin/env bash

# Ignore Swagger Generated Code
#
# There have been instances where the gosec tool will mark the generated
# code as containing a hard-coded password and thus fail the build due to
# long complicated resource names in the swagger such as: 
# `coolnewservice_accounts`
# This script prevents this by ignoring the code generated by our swagger
# documentation.

platform=$(uname)

## Generate Client
rm -Rf ./internal/swagger-client/interview-accountapi
mkdir -p ./internal/swagger-client/interview-accountapi
swagger generate client -f ./swagger.yaml -t ./internal/swagger-client/interview-accountapi

for f in $(find ./internal/swagger-client -type f -name "*.go"); do
    if [[ "$platform" == "Darwin" ]]; then
      sed -i '' '/^\/\/ Code generated by/s/^/\/\* #nosec \*\//' ${f}
    else
      sed -i '/^\/\/ Code generated by/s/^/\/\* #nosec \*\//' ${f}
    fi
done


# Generate Service Models
rm -Rf ./internal/app/interview-accountapi/api/externalmodels
swagger generate model -f swagger.yaml
mv models ./internal/app/interview-accountapi/api/externalmodels

for f in $(find ./internal/app/interview-accountapi/api/externalmodels -type f -name "*.go"); do
    if [[ "$platform" == "Darwin" ]]; then
      sed -i '' '/^\/\/ Code generated by/s/^/\/\* #nosec \*\//' ${f}
    else
      sed -i '/^\/\/ Code generated by/s/^/\/\* #nosec \*\//' ${f}
    fi
done

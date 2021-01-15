#!/bin/bash

# exit the script on command errors or unset variables
# http://redsymbol.net/articles/unofficial-bash-strict-mode/
set -euo pipefail
IFS=$'\n\t'

readonly location="westus2"
# keyvaults must be uniquely named, so let's give it a random number
readonly rand=$((1 + RANDOM % 1000))
readonly rg_name="soft-delete-demo"

# print command before running it
set -x

az group create \
    --location "$location" \
    --name "$rg_name"

az keyvault create \
    --location "$location" \
    --name "with-soft-delete-$rand" \
    --resource-group "$rg_name" \
    --enable-soft-delete true \
    --sku standard

# turn off command printing
{ set +x; } 2>/dev/null

echo "Created keyvaults:"
echo "- with-soft-delete-$rand"
echo "next step: go run . with-soft-delete-$rand"

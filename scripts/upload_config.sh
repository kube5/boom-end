#!/bin/sh

# TODO: support consul address as argument

SCRIPT_DIR=$(dirname $0)
CONFIG_FILE=$SCRIPT_DIR/../config.json

KEY_TO_PUT=$1

keys_array=($(jq -r 'to_entries | .[] | "\(.key)"' "$CONFIG_FILE"))

for index in "${!keys_array[@]}"; do
    key="${keys_array[$index]}"
    value=$(jq -c '.'$key $CONFIG_FILE)

    if [ "$key" == "$KEY_TO_PUT" ] || [ "$KEY_TO_PUT" == "" ]; then
        echo "Setting $key to $value"
        consul kv put mu/config/$key "${value}"
    fi
done

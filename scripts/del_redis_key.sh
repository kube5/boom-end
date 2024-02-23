#!/bin/bash

# Redis connection information
REDIS_HOST="localhost"
REDIS_PORT="6379"
REDIS_PASSWORD=""  # If there is a password, provide the corresponding password

# Get the keyword to be deleted from command-line arguments
KEYWORD="$1"

# Check if the keyword is empty
if [[ -z "$KEYWORD" ]]; then
  echo "Please provide the keyword to be deleted as an argument!"
  exit 1
fi

# Use the redis-cli command to connect to Redis and execute the KEYS command to get a list of keys containing the keyword
KEYS=$(redis-cli -h $REDIS_HOST -p $REDIS_PORT  KEYS "*$KEYWORD*")

# Iterate over the key list and use the DEL command to delete each key
for KEY in $KEYS
do
    echo "Deleting Key: $KEY"
    redis-cli -h $REDIS_HOST -p $REDIS_PORT  DEL $KEY
done

echo "Deletion completed."

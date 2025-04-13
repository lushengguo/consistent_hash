#!/bin/bash

# Read server addresses from setting.json
SERVERS=$(jq -r '.servers[]' setting.json)

# Launch each server in the background
for SERVER in $SERVERS; do
  SERVER_ADDRESS=$SERVER cargo run --bin server &
done

echo "All servers are running."

# Wait for all background processes
wait
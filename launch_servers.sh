#!/bin/bash

# Read server addresses from setting.json
SERVERS=$(jq -r '.servers[]' setting.json)

# Launch each server in the background
for SERVER in $SERVERS; do
  cargo run --bin server "$SERVER" &
done

echo "All servers are running."

# Wait for all background processes
wait
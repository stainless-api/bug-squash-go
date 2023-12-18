#!/usr/bin/env bash

set -euo pipefail

root_dir=$(dirname "$(dirname "$0")")

prism_cli="$root_dir/prism/node_modules/.bin/prism"

kill_server_on_port() {
  pids=$(lsof -t -i tcp:"$1" || echo "")
  if [ "$pids" != "" ]; then
    kill "$pids"
    echo "Stopped $pids."
  fi
}

# Ensure cleanup is called on script exit, error or interrupt
cleanup() {
  echo
  echo "Stopping servers..."
  kill_server_on_port "8077" || true
  kill_server_on_port "4010" || true
}
trap cleanup EXIT

# Build project
go build ./...

# Clear any existing logs and pids
rm -f "*.log"

# Start integration mock server in background
echo "Starting integration mock server..."
kill_server_on_port "8077"
(go run scripts/integration_server.go > .integration.log 2> .integration-err.log &)

# Ensure we have node_modules
if [ ! -d "$root_dir/prism/node_modules" ]; then
  echo "Installing prism mock server..."
  cd "$root_dir/prism"
  npm install
  cd -
fi

# Start prism mock server in background
kill_server_on_port "4010"
echo -n "Starting prism mock server..."
($prism_cli mock openapi.json > .prism.log 2> .prism-err.log &)

# Wait til prism has started
while ! grep -q "Prism is listening" ".prism.log" ; do
  echo -n "."
  sleep 0.1
done
echo

# Run test suite
echo "Running tests..."
go test ./...

# Normal exit
echo "Tests completed successfully. ✨"

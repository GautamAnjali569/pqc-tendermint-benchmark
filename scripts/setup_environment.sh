
#!/bin/bash

# Setup script for pqc-tendermint-benchmark

echo "Setting up the benchmarking environment..."

# Update and install dependencies
sudo apt update && sudo apt install -y golang

# Set Go environment variables
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# Install required Go packages
go install github.com/cloudflare/circl

echo "Environment setup complete. You can now run the benchmarks."

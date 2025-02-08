#!/bin/bash

# Run benchmark script for pqc-tendermint-benchmark

echo "Running benchmarking scripts..."

# Execute benchmarks and redirect output to results file
go run main.go | tee results/results.log

echo "Benchmarking completed. Check results/results.log for details."



package main

import (
	"fmt"
	"os"
	"pqc-tendermint-benchmark/benchmarks"
)

func main() {
	fmt.Println("Starting Dilithium vs. Ed25519 Benchmarking...")
	
	benchmarks.RunBenchmarks()
	
	fmt.Println("Benchmarking Completed. Check results directory for logs.")
}

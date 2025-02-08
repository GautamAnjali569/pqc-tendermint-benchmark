
package benchmarks

import (
	"fmt"
	"time"
	"github.com/cloudflare/circl/sign/dilithium"
)

// RunDilithiumBenchmark executes benchmarks for the Dilithium signature algorithm.
func RunDilithiumBenchmark() {
	fmt.Println("Running Dilithium Benchmark...")
	
	// Key Generation Benchmark
	start := time.Now()
	_, pk, _ := dilithium.Mode5.GenerateKey(nil)
	keyGenTime := time.Since(start)
	fmt.Printf("Key Generation Time: %v\n", keyGenTime)
	
	// Signing Benchmark
	message := []byte("Benchmark Test Message")
	start = time.Now()
	sk := dilithium.Mode5.NewKeyFromSeed(pk.Seed())
	sig := sk.Sign(message)
	signTime := time.Since(start)
	fmt.Printf("Signing Time: %v\n", signTime)
	
	// Verification Benchmark
	start = time.Now()
	valid := pk.Verify(message, sig)
	verifyTime := time.Since(start)
	fmt.Printf("Verification Time: %v\n", verifyTime)
	fmt.Printf("Signature Valid: %v\n", valid)
}

// RunBenchmarks executes all signature benchmarks.
func RunBenchmarks() {
	RunDilithiumBenchmark()
}

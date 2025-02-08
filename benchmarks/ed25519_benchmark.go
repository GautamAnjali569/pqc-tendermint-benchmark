package benchmarks

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
	"time"
)

// RunEd25519Benchmark executes benchmarks for the Ed25519 signature algorithm.
func RunEd25519Benchmark() {
	fmt.Println("Running Ed25519 Benchmark...")
	
	// Key Generation Benchmark
	start := time.Now()
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)
	keyGenTime := time.Since(start)
	fmt.Printf("Key Generation Time: %v\n", keyGenTime)
	
	// Signing Benchmark
	message := []byte("Benchmark Test Message")
	start = time.Now()
	signature := ed25519.Sign(privateKey, message)
	signTime := time.Since(start)
	fmt.Printf("Signing Time: %v\n", signTime)
	
	// Verification Benchmark
	start = time.Now()
	valid := ed25519.Verify(publicKey, message, signature)
	verifyTime := time.Since(start)
	fmt.Printf("Verification Time: %v\n", verifyTime)
	fmt.Printf("Signature Valid: %v\n", valid)
}

// RunBenchmarks executes all signature benchmarks.
func RunBenchmarks() {
	RunDilithiumBenchmark()
	RunEd25519Benchmark()
}

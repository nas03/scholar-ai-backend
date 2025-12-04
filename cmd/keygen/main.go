package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nas03/scholar-ai/backend/global"
	"github.com/nas03/scholar-ai/backend/internal/helper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Initialize logger for keygen command
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05")
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.DisableStacktrace = true
	logger, err := cfg.Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	global.Log = logger
	defer func() {
		_ = global.Log.Sync()
	}()

	// Parse command line flags
	keyPath := flag.String("key", "keys/private_key.pem", "Path to output private key file")
	certPath := flag.String("cert", "keys/certificate.pem", "Path to output certificate file (optional, empty to skip)")
	flag.Parse()

	// Create keys directory if it doesn't exist
	keyDir := filepath.Dir(*keyPath)
	if err := os.MkdirAll(keyDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating keys directory: %v\n", err)
		os.Exit(1)
	}

	// Check if key file already exists
	if _, err := os.Stat(*keyPath); err == nil {
		fmt.Fprintf(os.Stderr, "Error: Private key file already exists at %s\n", *keyPath)
		fmt.Fprintf(os.Stderr, "To regenerate, please delete the existing file first.\n")
		os.Exit(1)
	}

	// Step 1: Generate RSA private key (2048-bit is standard for RS256)
	fmt.Println("Generating RSA private key (2048-bit)...")
	privateKey, err := helper.CreatePrivateKey(*keyPath, *certPath)
	if err != nil {
		global.Log.Error("Error creating a private key", zap.Error(err))
		os.Exit(1)
	}

	fmt.Printf("\n⚠️  IMPORTANT: Keep your private key secure and never commit it to version control!\n")

	// Display public key info
	fmt.Printf("\nPublic Key Info:\n")
	fmt.Printf("  Modulus size: %d bits\n", privateKey.N.BitLen())
	fmt.Printf("  Public exponent: %d\n", privateKey.PublicKey.E)
}

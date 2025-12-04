package helper

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
	"time"

	"github.com/nas03/scholar-ai/backend/global"
	"go.uber.org/zap"
)

// CreatePrivateKey generates a new RSA private key (2048-bit for RS256)
func CreatePrivateKey(keyPath, certPath string) (*rsa.PrivateKey, error) {
	// Step 1: Generate RSA private key (2048-bit is standard for RS256)
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		global.Log.Error("Error generating RSA private key", zap.Error(err))
		return nil, err
	}

	// Step 2: Create X.509 certificate template
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization:  []string{"Scholar AI"},
			Country:       []string{"VN"},
			Province:      []string{""},
			Locality:      []string{""},
			StreetAddress: []string{""},
			PostalCode:    []string{""},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour), // Valid for 1 year
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// Step 3: Create certificate (self-signed for development)
	// In production, you'd use a Certificate Authority (CA)
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		global.Log.Error("Error creating certificate", zap.Error(err))
		return nil, err
	}

	// Step 4: Encode private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// Step 5: Write private key to file
	privateKeyFile, err := os.OpenFile(keyPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		global.Log.Error("Error creating private key file", zap.Error(err), zap.String("path", keyPath))
		return nil, err
	}
	defer privateKeyFile.Close()

	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		global.Log.Error("Error encoding private key to PEM", zap.Error(err))
		return nil, err
	}

	// Ensure secure file permissions (0600 = read/write for owner only)
	if err := os.Chmod(keyPath, 0600); err != nil {
		global.Log.Warn("Failed to set secure permissions on private key", zap.Error(err), zap.String("path", keyPath))
	}

	global.Log.Info("Private key saved successfully", zap.String("path", keyPath))

	// Step 6: Encode certificate to PEM format (optional, but useful)
	if certPath != "" {
		certPEM := &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certDER,
		}

		certFile, err := os.OpenFile(certPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
		if err != nil {
			global.Log.Error("Error creating certificate file", zap.Error(err), zap.String("path", certPath))
			return nil, err
		}
		defer certFile.Close()

		if err := pem.Encode(certFile, certPEM); err != nil {
			global.Log.Error("Error encoding certificate to PEM", zap.Error(err))
			return nil, err
		}

		// Set file permissions (certificate can be readable by all)
		if err := os.Chmod(certPath, 0644); err != nil {
			global.Log.Warn("Failed to set permissions on certificate", zap.Error(err), zap.String("path", certPath))
		}

		global.Log.Info("Certificate saved successfully", zap.String("path", certPath))
	}

	return privateKey, nil
}

// GetPublicKey extracts the public key from a private key
func GetPublicKey(keyPath string) (*rsa.PublicKey, error) {
	// Read the PEM file
	if keyPath == "" {
		keyPath = "keys/private_key.pem"
	}
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		global.Log.Error("Error reading private key file", zap.Error(err), zap.String("path", keyPath))
		return nil, err
	}

	// Decode PEM block
	block, _ := pem.Decode(keyData)
	if block == nil {
		global.Log.Error("Error decoding PEM block")
		return nil, errors.New("failed to decode PEM block: invalid PEM format")
	}

	// Parse the private key
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		global.Log.Error("Error parsing private key", zap.Error(err))
		return nil, err
	}

	return &privateKey.PublicKey, nil
}

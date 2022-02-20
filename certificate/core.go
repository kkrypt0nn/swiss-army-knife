package certificate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"time"
)

// GetLocalIP returns the non loopback local IP of the host.
func GetLocalIP() net.IP {
	connection, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil
	}
	defer connection.Close()
	localIP := connection.LocalAddr().(*net.UDPAddr)
	return localIP.IP
}

// GenerateCertificate generates a self-signed certificate for the given host and returns the private RSA key and the generated certificate PEM encoded.
func GenerateCertificate(serial int64, country string, organization string, commonName string, expire time.Time, localIP net.IP) ([]byte, []byte) {
	certificateTemplate := &x509.Certificate{
		IsCA:                  true,
		BasicConstraintsValid: true,
		SerialNumber:          big.NewInt(serial),
		Subject: pkix.Name{
			Country:      []string{country},
			Organization: []string{organization},
			CommonName:   commonName,
		},
		NotBefore: time.Now(),
		NotAfter:  expire,
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth,
		},
	}
	if localIP != nil {
		certificateTemplate.IPAddresses = append(certificateTemplate.IPAddresses, localIP)
	} else {
		certificateTemplate.IPAddresses = append(certificateTemplate.IPAddresses, GetLocalIP())
	}
	certificateTemplate.IPAddresses = append(certificateTemplate.IPAddresses, net.ParseIP("127.0.0.1"))
	certificateTemplate.IPAddresses = append(certificateTemplate.IPAddresses, net.ParseIP("::1"))
	certificateTemplate.DNSNames = append(certificateTemplate.DNSNames, "localhost")

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil
	}

	certificate, err := x509.CreateCertificate(rand.Reader, certificateTemplate, certificateTemplate, &privKey.PublicKey, privKey)
	if err != nil {
		return nil, nil
	}

	privateKey := x509.MarshalPKCS1PrivateKey(privKey)
	privateKey = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKey,
	})

	certificate = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certificate,
	})

	return privateKey, certificate
}

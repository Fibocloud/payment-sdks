package tdb_cg

import (
	"crypto/tls"
	"crypto/x509"
	"os"

	"golang.org/x/crypto/pkcs12"
)

// func (t *tdbcg) CertParser() {
// 	// certPathPfx := "/path/to/certificate.pfx"  // Replace with your actual PFX certificate path
// 	// certPass := "certificate_password"             // Replace with your actual certificate password
// 	// certPathCer := "/path/to/root_certificate.cer" // Replace with your actual CER certificate path

// 	// Load the PFX certificate
// 	pfxData, err := os.ReadFile(t.certPathPfx)
// 	if err != nil {
// 		fmt.Println("Error reading PFX certificate file:", err)
// 		return
// 	}

// 	// Load the CER certificate
// 	cerData, err := os.ReadFile(t.certPathCer)
// 	if err != nil {
// 		fmt.Println("Error reading CER certificate file:", err)
// 		return
// 	}

// 	// Parse the PFX certificate
// 	_, cert, err := pkcs12.Decode(pfxData, t.certPass)
// 	if err != nil {
// 		fmt.Println("Error parsing PFX certificate:", err)
// 		return
// 	}

// 	// Parse the CER certificate
// 	rootCert, err := x509.ParseCertificate(cerData)
// 	if err != nil {
// 		fmt.Println("Error parsing CER certificate:", err)
// 		return
// 	}

// 	// Create a certificate pool with the root certificate
// 	roots := x509.NewCertPool()
// 	roots.AddCert(rootCert)

// 	// Create an X.509 chain
// 	intermediates := x509.NewCertPool()
// 	opts := x509.VerifyOptions{
// 		Roots:         roots,
// 		Intermediates: intermediates,
// 	}

// 	chains, err := cert.Verify(opts)
// 	if err != nil {
// 		fmt.Println("Certificate verification failed:", err)
// 		return
// 	}

// 	// Certificate and root certificate are valid
// 	fmt.Println("Certificate and root certificate are valid.")
// }

func LoadP12TLSCfg(pfxPath, password string) (*x509.CertPool, tls.Certificate, error) {
	data, err := os.ReadFile(pfxPath)
	if err != nil {
		return nil, tls.Certificate{}, err
	}

	// certs, err := ioutil.ReadFile(certPath)
	// if err != nil {
	// 	return nil, tls.Certificate{}, err
	// }

	pk, crt, err := pkcs12.Decode(data, password)
	if err != nil {
		return nil, tls.Certificate{}, err
	}
	pool := x509.NewCertPool()
	pool.AddCert(crt)
	tlsCrt := tls.Certificate{
		Certificate: [][]byte{crt.Raw},
		PrivateKey:  pk,
		Leaf:        crt,
	}
	return pool, tlsCrt, nil
}

func LoadClientTLSCfg(pfxPath, password, serverName string) (*tls.Config, error) {
	pool, crt, err := LoadP12TLSCfg(pfxPath, password)
	if err != nil {
		return nil, err
	}
	cfg := &tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{crt},
		ServerName:   serverName,
	}
	return cfg, nil
}

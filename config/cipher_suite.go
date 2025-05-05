package config

import (
	"crypto/tls"
	"fmt"
)

// ParseCipherSuite converts a cipher suite name to its corresponding constant.
func ParseCipherSuite(cipher string) (uint16, error) {
	switch cipher {
	case "TLS_RSA_WITH_AES_128_CBC_SHA":
		return tls.TLS_RSA_WITH_AES_128_CBC_SHA, nil
	case "TLS_RSA_WITH_AES_256_CBC_SHA":
		return tls.TLS_RSA_WITH_AES_256_CBC_SHA, nil
	case "TLS_RSA_WITH_AES_128_GCM_SHA256":
		return tls.TLS_RSA_WITH_AES_128_GCM_SHA256, nil
	case "TLS_RSA_WITH_AES_256_GCM_SHA384":
		return tls.TLS_RSA_WITH_AES_256_GCM_SHA384, nil
	case "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":
		return tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA, nil
	case "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":
		return tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA, nil
	case "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":
		return tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, nil
	case "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384":
		return tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, nil
	case "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256":
		return tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256, nil
	case "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384":
		return tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, nil
	case "TLS_AES_128_GCM_SHA256":
		return tls.TLS_AES_128_GCM_SHA256, nil
	case "TLS_AES_256_GCM_SHA384":
		return tls.TLS_AES_256_GCM_SHA384, nil
	case "TLS_CHACHA20_POLY1305_SHA256":
		return tls.TLS_CHACHA20_POLY1305_SHA256, nil
	default:
		return 0, fmt.Errorf("unsupported cipher suite: %s", cipher)
	}
}

package gateway

import (
	"fmt"
	"sigs.k8s.io/aws-load-balancer-controller/test/framework/utils"
)

// Helper function to generate random OIDC credentials
func GenerateOIDCCredentials() (clientID string, clientSecret string) {
	// Generate random 20-character alphanumeric client ID
	clientID = fmt.Sprintf("test-client-%s", utils.RandomDNS1123Label(12))

	// Generate random 32-character base64-like client secret
	clientSecret = fmt.Sprintf("%s%s", utils.RandomDNS1123Label(16), utils.RandomDNS1123Label(16))

	return clientID, clientSecret
}

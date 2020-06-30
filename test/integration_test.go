package test

import (
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration test")
	}

	t.Logf("Kubeconfig: %v", os.Getenv("KUBECONFIG"))
}

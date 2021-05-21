// +build integration

package tests

import (
	"os"
	"testing"

	"github.com/jetstack/cert-manager/test/acme/dns"

	pinto "gitlab.com/whizus/customer/pinto/cert-manager-webhook-pinto/pkg/dns"
)

var (
	zone = os.Getenv("TEST_ZONE_NAME")
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getting current working dir: %s", err.Error())
	}

	fixture := dns.NewFixture(&pinto.ProviderSolver{},
		dns.SetResolvedZone(zone),
		dns.SetAllowAmbientCredentials(true),
		dns.SetBinariesPath(currentDir+"/kubebuilder/bin"),
		dns.SetManifestPath(currentDir+"/testdata"),
		dns.SetStrict(true),
	)

	fixture.RunConformance(t)
}

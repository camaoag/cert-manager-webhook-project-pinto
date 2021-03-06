//go:build integration
// +build integration

package tests

import (
	"os"
	"testing"

	"github.com/jetstack/cert-manager/test/acme/dns"

	pinto "github.com/camaoag/cert-manager-webhook-project-pinto/pkg/dns"
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
		dns.SetDNSServer("ns3.digitalocean.com:53"),
		dns.SetResolvedZone(zone),
		dns.SetAllowAmbientCredentials(true),
		//dns.SetBinariesPath(),
		dns.SetManifestPath(currentDir+"/testdata"),
		dns.SetStrict(true),
	)

	fixture.RunConformance(t)
}

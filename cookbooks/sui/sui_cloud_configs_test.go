package sui_cookbooks

import (
	"context"
	"fmt"
)

func (t *SuiCookbookTestSuite) testBuildAndUpload(cloudProvider, network string) {
	cfg := SuiConfigOpts{
		WithLocalNvme:        true,
		DownloadSnapshot:     false,
		WithIngress:          true,
		WithServiceMonitor:   true,
		WithArchivalFallback: true,
		CloudProvider:        cloudProvider,
		Network:              network,
	}
	suiNodeDefinition = GetSuiClientClusterDef(cfg)
	//t.TestCreateClusterClass()
	// ^ only needed if no pre-existing cluster definition
	t.Require().Equal(fmt.Sprintf("sui-%s-%s", network, cloudProvider), suiNodeDefinition.ClusterClassName)
	_, err := suiNodeDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)
}

func (t *SuiCookbookTestSuite) TestNvmeConfigs() {
	cps := []string{"aws", "gcp", "do"}
	networks := []string{"mainnet", "testnet", "devnet"}
	for _, cp := range cps {
		for _, network := range networks {
			t.testBuildAndUpload(cp, network)
		}
	}
}

func (t *SuiCookbookTestSuite) TestAwsGp3Config() {
	cps := []string{"aws"}
	networks := []string{"devnet"}
	for _, cp := range cps {
		for _, network := range networks {
			t.testBuildAndUploadEbs(cp, network)
		}
	}
}

func (t *SuiCookbookTestSuite) testBuildAndUploadEbs(cloudProvider, network string) {
	cfg := SuiConfigOpts{
		WithLocalNvme:        false,
		DownloadSnapshot:     true,
		WithIngress:          false,
		WithServiceMonitor:   true,
		WithArchivalFallback: false,
		CloudProvider:        cloudProvider,
		Network:              network,
	}
	suiNodeDefinition = GetSuiClientClusterDef(cfg)
	//t.TestCreateClusterClass()
	// ^ only needed if no pre-existing cluster definition
	t.Require().Equal(fmt.Sprintf("sui-%s-%s-ssd", network, cloudProvider), suiNodeDefinition.ClusterClassName)

	gcd := suiNodeDefinition.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	err := gcd.CreateClusterClassDefinitions(context.Background(), t.ZeusTestClient)
	t.Require().Nil(err)
	_, err = suiNodeDefinition.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)
}

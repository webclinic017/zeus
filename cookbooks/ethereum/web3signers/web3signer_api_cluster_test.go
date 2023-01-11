package web3signer_cookbooks

import (
	"context"

	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_resp_types/topology_workloads"
	zeus_topology_config_drivers "github.com/zeus-fyi/zeus/pkg/zeus/config_drivers"
)

func (t *Web3SignerCookbookTestSuite) TestClusterAPIDeploy() {
	t.TestUploadWeb3SignerAPIChart()
	ctx := context.Background()
	resp, err := t.ZeusTestClient.DeployCluster(ctx, Web3SignerExternalAPIClusterDefinition)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *Web3SignerCookbookTestSuite) TestClusterAPIDestroy() {
	ctx := context.Background()
	knsReq := DeployWeb3SignerExternalAPIKnsReq
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, knsReq)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *Web3SignerCookbookTestSuite) TestCreateAPIClusterBase() {
	ctx := context.Background()
	basesInsert := []string{Web3SignerExternalAPIClusterBaseName, Web3SignerExternalAPIClusterIngressBaseName}
	cc := zeus_req_types.TopologyCreateOrAddComponentBasesToClassesRequest{
		ClusterClassName:   Web3SignerExternalAPIClusterClassName,
		ComponentBaseNames: basesInsert,
	}
	_, err := t.ZeusTestClient.AddComponentBasesToClass(ctx, cc)
	t.Require().Nil(err)
}

func (t *Web3SignerCookbookTestSuite) TestCreateAPIClusterSkeletonBases() {
	ctx := context.Background()

	cc := zeus_req_types.TopologyCreateOrAddSkeletonBasesToClassesRequest{
		ClusterClassName:  Web3SignerExternalAPIClusterClassName,
		ComponentBaseName: Web3SignerExternalAPIClusterBaseName,
		SkeletonBaseNames: []string{Web3SignerExternalAPIClusterSkeletonBaseName, Web3SignerExternalAPIClusterIngressSkeletonBaseName},
	}
	_, err := t.ZeusTestClient.AddSkeletonBasesToClass(ctx, cc)
	t.Require().Nil(err)

}
func (t *Web3SignerCookbookTestSuite) TestUploadWeb3SignerAPIChart() {
	ctx := context.Background()

	inf := topology_workloads.NewTopologyBaseInfraWorkload()
	err := Web3SignerChartPath.WalkAndApplyFuncToFileType(".yaml", inf.DecodeK8sWorkload)
	t.Require().Nil(err)

	ingr := topology_workloads.NewTopologyBaseInfraWorkload()
	err = Web3SignerIngressChartPath.WalkAndApplyFuncToFileType(".yaml", ingr.DecodeK8sWorkload)
	t.Require().Nil(err)

	inf.Ingress = ingr.Ingress

	infCfg := zeus_topology_config_drivers.IngressDriver{NginxAuthURL: t.AuthURL}
	stsCfg := GetWeb3SignerAPIStatefulSetConfig(t.CustomWeb3SignerImage)
	tc := zeus_topology_config_drivers.TopologyConfigDriver{
		IngressDriver:     &infCfg,
		StatefulSetDriver: &stsCfg,
	}

	tc.SetCustomConfig(&inf)

	resp, err := t.ZeusTestClient.UploadChart(ctx, Web3SignerChartPath, Web3SignerChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployWeb3SignerExternalAPIKnsReq.TopologyID = resp.TopologyID
	tar := zeus_req_types.TopologyRequest{TopologyID: DeployWeb3SignerExternalAPIKnsReq.TopologyID}
	chartResp, err := t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(Web3SignerChartPath)
	t.Require().Nil(err)

	resp, err = t.ZeusTestClient.UploadChart(ctx, Web3SignerIngressChartPath, Web3SignerIngressChart)
	t.Require().Nil(err)
	t.Assert().NotZero(resp.TopologyID)

	DeployWeb3SignerExternalAPIKnsReq.TopologyID = resp.TopologyID
	tar = zeus_req_types.TopologyRequest{TopologyID: DeployWeb3SignerExternalAPIKnsReq.TopologyID}
	chartResp, err = t.ZeusTestClient.ReadChart(ctx, tar)
	t.Require().Nil(err)
	t.Assert().NotEmpty(chartResp)

	err = chartResp.PrintWorkload(Web3SignerIngressChartPath)
	t.Require().Nil(err)
}

package avax_node_cookbooks

import (
	"fmt"

	zk8s_clusters "github.com/zeus-fyi/zeus/zeus/z_client/clusters"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
)

func (t *AvaxCookbookTestSuite) TestFujiClusterDeploy() {
	cd := AvaxNodeClusterDefinition
	cd.ClusterClassName = "avaxFujiNode"
	cd.CloudCtxNs.Namespace = "avax-fuji"
	delete(cd.ComponentBases, "avaxIngress")
	cd.ComponentBases["avaxClients"] = ConfigureAvaxNodeClusterBase(ctx, "fuji")
	_, err := cd.UploadChartsFromClusterDefinition(ctx, t.ZeusTestClient, true)
	t.Require().Nil(err)
	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)

	cdep := cd.GenerateDeploymentRequest()
	_, err = zk8s_clusters.DeployCluster(ctx, t.ZeusTestClient, cdep)
	t.Require().Nil(err)

}

func (t *AvaxCookbookTestSuite) TestFujiClusterDestroy() {
	d := zeus_req_types.TopologyDeployRequest{
		CloudCtxNs: AvaxNodeCloudCtxNs,
	}
	d.Namespace = "avax-fuji"
	resp, err := t.ZeusTestClient.DestroyDeploy(ctx, d)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

func (t *AvaxCookbookTestSuite) TestFujiClusterSetup() {
	cd := AvaxNodeClusterDefinition
	cd.ClusterClassName = "avaxFujiNode"
	cd.CloudCtxNs.Namespace = "avax-fuji"
	delete(cd.ComponentBases, "avaxIngress")
	cd.ComponentBases["avaxClients"] = ConfigureAvaxNodeClusterBase(ctx, "fuji")
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)

	gdr := cd.GenerateDeploymentRequest()
	t.Assert().NotEmpty(gdr)
	fmt.Println(gdr)

	sbDefs, err := cd.GenerateSkeletonBaseCharts()
	t.Require().Nil(err)
	t.Assert().NotEmpty(sbDefs)
}

func (t *AvaxCookbookTestSuite) TestClusterFujiDefinitionCreation() {
	cd := AvaxNodeClusterDefinition
	cd.ClusterClassName = "avaxFujiNode"
	cd.CloudCtxNs.Namespace = "avax-fuji"
	delete(cd.ComponentBases, "avaxIngress")
	cd.ComponentBases["avaxClients"] = ConfigureAvaxNodeClusterBase(ctx, "fuji")
	gcd := cd.BuildClusterDefinitions()
	t.Assert().NotEmpty(gcd)
	fmt.Println(gcd)
	err := gcd.CreateClusterClassDefinitions(ctx, t.ZeusTestClient)
	t.Require().Nil(err)
}

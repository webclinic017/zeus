package pods_client

import (
	"context"

	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	zeus_pods_reqs "github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types/pods"
)

var ctx = context.Background()

func (t *PodsClientTestSuite) TestPortForwardReqToPods() {
	deployKnsReq.TopologyID = 0

	cliReq := zeus_pods_reqs.ClientRequest{
		MethodHTTP: "GET",
		Endpoint:   "eth/v1/node/syncing",
		Ports:      []string{"5052:5052"},
	}
	filter := strings_filter.FilterOpts{DoesNotInclude: []string{"geth"}}

	par := zeus_pods_reqs.PodActionRequest{
		TopologyDeployRequest: deployKnsReq,
		Action:                zeus_pods_reqs.PortForwardToAllMatchingPods,
		PodName:               "zeus-lighthouse-0",
		ContainerName:         "lighthouse",
		ClientReq:             &cliReq,
		FilterOpts:            &filter,
	}

	resp, err := t.ZeusTestClient.PortForwardReqToPods(ctx, par)
	t.Require().Nil(err)
	t.Assert().NotEmpty(resp)
}

package zeus_client

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"

	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func (z *ZeusClient) DeployCluster(ctx context.Context, tar zeus_req_types.ClusterTopologyDeployRequest) (zeus_resp_types.ClusterStatus, error) {
	z.PrintReqJson(tar)
	respJson := zeus_resp_types.ClusterStatus{}
	resp, err := z.R().
		SetResult(&respJson).
		SetBody(tar).
		Post(zeus_endpoints.DeployClusterTopologyV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		if resp.StatusCode() == http.StatusBadRequest {
			err = errors.New("bad request")
		}
		log.Err(err).Msg("ZeusClient: DeployCluster")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

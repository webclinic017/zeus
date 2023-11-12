package zeus_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/compression"
	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	zeus_endpoints "github.com/zeus-fyi/zeus/zeus/z_client/endpoints"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_req_types"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_resp_types"
)

func (z *ZeusClient) UploadChart(ctx context.Context, p filepaths.Path, tar zeus_req_types.TopologyCreateRequest) (zeus_resp_types.TopologyCreateResponse, error) {
	respJson := zeus_resp_types.TopologyCreateResponse{}
	err := z.ZipK8sChartToPath(&p)
	if err != nil {
		return respJson, err
	}
	z.PrintReqJson(tar)
	resp, err := z.R().
		SetResult(&respJson).
		SetFormData(map[string]string{
			"topologyName":      tar.TopologyName,
			"chartName":         tar.ChartName,
			"chartDescription":  tar.ChartDescription,
			"version":           tar.Version,
			"clusterClassName":  tar.ClusterClassName,
			"componentBaseName": tar.ComponentBaseName,
			"skeletonBaseName":  tar.SkeletonBaseName,
			"tag":               tar.Tag,
		}).
		SetFile("chart", p.FileOutPath()).
		Post(zeus_endpoints.InfraCreateV1Path)

	if err != nil || resp.StatusCode() >= 400 {
		if err == nil {
			err = fmt.Errorf("non-OK status code: %d", resp.StatusCode())
		}
		log.Err(err).Msg("ZeusClient: UploadChart")
		return respJson, err
	}
	z.PrintRespJson(resp.Body())
	return respJson, err
}

func (z *ZeusClient) ZipK8sChartToPath(p *filepaths.Path) error {
	comp := compression.NewCompression()
	err := comp.GzipCompressDir(p)
	if err != nil {
		log.Err(err).Interface("path", p).Msg("ZeusClient: ZipK8sChartToPath")
		return err
	}
	return err
}

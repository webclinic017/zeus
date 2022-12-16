package beacon_cookbooks

import (
	"fmt"
	"time"

	filepaths "github.com/zeus-fyi/zeus/pkg/utils/file_io/lib/v0/paths"
	strings_filter "github.com/zeus-fyi/zeus/pkg/utils/strings"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_common_types"
	"github.com/zeus-fyi/zeus/pkg/zeus/client/zeus_req_types"
)

const (
	start    = "start"
	download = "download"
)

var BeaconCloudCtxNs = zeus_common_types.CloudCtxNs{
	CloudProvider: "do",
	Region:        "sfo3",
	Context:       "do-sfo3-dev-do-sfo3-zeus",
	Namespace:     "beacon", // set with your own namespace
	Env:           "production",
}

// chart workload metadata
var ingressChart = zeus_req_types.TopologyCreateRequest{
	TopologyName:      "beaconIngress",
	ChartName:         "beaconIngress",
	ChartDescription:  "beaconIngress",
	Version:           fmt.Sprintf("beaconIngress-v.0.%d", time.Now().Unix()),
	SkeletonBaseName:  "beaconIngress",
	ComponentBaseName: "beaconIngress",
	ClusterBaseName:   "ethereumBeacon",
	Tag:               "latest",
}

var ingressChartPath = filepaths.Path{
	PackageName: "",
	DirIn:       "./ethereum/beacons/infra/ingress",
	DirOut:      "./ethereum/beacons/infra/processed_beacon_ingress",
	FnIn:        "beaconIngress", // filename for your gzip workload
	FnOut:       "",
	Env:         "",
	FilterFiles: strings_filter.FilterOpts{},
}

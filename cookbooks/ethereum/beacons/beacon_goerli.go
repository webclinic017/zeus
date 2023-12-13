package ethereum_beacon_cookbooks

import (
	zeus_cluster_config_drivers "github.com/zeus-fyi/zeus/zeus/cluster_config_drivers"
	"github.com/zeus-fyi/zeus/zeus/z_client/zeus_common_types"
)

var (
	BeaconGoerliClusterDefinition = zeus_cluster_config_drivers.ClusterDefinition{
		ClusterClassName: "ethereum-goerli-beacons",
		CloudCtxNs:       BeaconGoerliCloudCtxNs,
		ComponentBases:   BeaconGoerliComponentBases,
	}
	BeaconGoerliCloudCtxNs = zeus_common_types.CloudCtxNs{
		CloudProvider: "do",
		Region:        "sfo3",
		Context:       "do-sfo3-dev-do-sfo3-zeus",
		Namespace:     "goerli-beacon", // set with your own namespace
		Env:           "production",
	}
	BeaconGoerliComponentBases = map[string]zeus_cluster_config_drivers.ComponentBaseDefinition{
		"consensus-clients": ConsensusClientGoerliComponentBase,
		"exec-clients":      ExecClientGoerliComponentBase,
	}
	ConsensusClientGoerliComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"lodestar-hercules": ConsensusClientGoerliSkeletonBaseConfig,
		},
	}
	ExecClientGoerliComponentBase = zeus_cluster_config_drivers.ComponentBaseDefinition{
		SkeletonBases: map[string]zeus_cluster_config_drivers.ClusterSkeletonBaseDefinition{
			"geth-hercules": ExecClientGoerliSkeletonBaseConfig,
		},
	}
)

package web3signer_cookbooks

import (
	web3signer_cmds_ai_generated "github.com/zeus-fyi/zeus/cookbooks/ethereum/web3signers/web3signer_cmds/ai_generated"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	v1 "k8s.io/api/core/v1"
)

func GetWeb3SignerAPIStatefulSetConfig(customImage string) config_overrides.StatefulSetDriver {
	args, _ := web3signer_cmds_ai_generated.Web3SignerAPICmd.CreateFieldsForCLI("eth2")
	port := v1.ContainerPort{
		Name:          "http",
		ContainerPort: 9000,
	}
	c := v1.Container{
		Name:      web3SignerClient,
		Image:     customImage,
		Command:   []string{"/bin/sh"},
		Args:      args,
		Ports:     []v1.ContainerPort{port},
		Env:       []v1.EnvVar{},
		Resources: v1.ResourceRequirements{},
	}
	contDriver := config_overrides.ContainerDriver{
		Container: c,
	}
	sc := config_overrides.StatefulSetDriver{}
	sc.ContainerDrivers = make(map[string]config_overrides.ContainerDriver)
	sc.ContainerDrivers[c.Name] = contDriver
	return sc
}

func GetWeb3SignerAPIServiceConfig() config_overrides.ServiceDriver {
	s := config_overrides.ServiceDriver{ExtendPorts: []v1.ServicePort{}}
	s.AddNginxTargetPort("nginx", "http")
	return s
}

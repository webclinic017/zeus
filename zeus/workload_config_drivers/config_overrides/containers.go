package config_overrides

import (
	v1Core "k8s.io/api/core/v1"
)

type ContainerDriver struct {
	IsDeleteContainer bool `json:"isDeleteContainer"`
	IsAppendContainer bool `json:"isAppendContainer"`
	IsInitContainer   bool `json:"isInitContainer"`
	v1Core.Container  `json:"v1Container"`
	AppendEnvVars     []v1Core.EnvVar `json:"appendEnvVars,omitempty"`
}

func (cd *ContainerDriver) SetContainerConfigs(cont *v1Core.Container) {
	if len(cd.Image) > 0 {
		cont.Image = cd.Image
	}
	if cd.Env != nil {
		cont.Env = cd.Env
	}
	if cd.AppendEnvVars != nil {
		cont.Env = append(cont.Env, cd.AppendEnvVars...)
	}
	if cd.Ports != nil {
		cont.Ports = cd.Ports
	}
	if cd.Command != nil {
		cont.Command = cd.Command
	}
	if cd.Args != nil {
		cont.Args = cd.Args
	}
	if cd.Resources.Limits != nil {
		cont.Resources.Limits = cd.Resources.Limits
	}
	if cd.Resources.Requests != nil {
		cont.Resources.Requests = cd.Resources.Requests
	}
	if cd.ImagePullPolicy != "" {
		cont.ImagePullPolicy = cd.ImagePullPolicy
	}

	if cd.VolumeMounts != nil {
		// if the driver has a matching name, then it will override the container's volume mount
		// otherwise, it will append to the container's volume mount
		m := make(map[string]v1Core.VolumeMount)
		for _, v := range cd.VolumeMounts {
			m[v.Name] = v
		}
		for i, v := range cont.VolumeMounts {
			if vm, ok := m[v.Name]; ok {
				cont.VolumeMounts[i] = vm
				delete(m, v.Name)
			}
		}
		for _, v := range m {
			cont.VolumeMounts = append(cont.VolumeMounts, v)
		}
	}
}

func (cd *ContainerDriver) CreateEnvVarKeyValue(k, v string) v1Core.EnvVar {
	return MakeKeyValueEnvVar(k, v)
}

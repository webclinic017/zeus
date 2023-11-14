package zk8s_templates

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/zeus/zeus/workload_config_drivers/config_overrides"
	v1 "k8s.io/api/apps/v1"
	v1Core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetStatefulSetTemplate(ctx context.Context, name string) *v1.StatefulSet {
	labels := GetLabels(ctx, name)
	selectors := GetSelector(ctx, name)
	return &v1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:   GetStatefulSetName(ctx, name),
			Labels: labels,
		},
		Spec: v1.StatefulSetSpec{
			Selector: metav1.SetAsLabelSelector(selectors),
			Template: v1Core.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1Core.PodSpec{},
			},
			ServiceName:         GetServiceName(ctx, name),
			PodManagementPolicy: v1.OrderedReadyPodManagement,
			UpdateStrategy: v1.StatefulSetUpdateStrategy{
				Type: v1.RollingUpdateStatefulSetStrategyType,
			},
		},
	}
}

func BuildStatefulSetDriver(ctx context.Context, containers Containers, sts StatefulSet) (config_overrides.StatefulSetDriver, error) {
	rc := int32(sts.ReplicaCount)
	stsDriver := config_overrides.StatefulSetDriver{
		ReplicaCount:     &rc,
		ContainerDrivers: make(map[string]config_overrides.ContainerDriver),
	}
	for containerName, container := range containers {
		contDriver, err := BuildContainerDriver(ctx, containerName, container)
		if err != nil {
			log.Error().Err(err).Msg("Failed to build container driver")
			return config_overrides.StatefulSetDriver{}, err
		}
		stsDriver.ContainerDrivers[containerName] = config_overrides.ContainerDriver{
			IsAppendContainer: true,
			IsInitContainer:   container.IsInitContainer,
			Container:         contDriver,
			AppendEnvVars:     nil,
		}
	}
	pvcCfg := config_overrides.PersistentVolumeClaimsConfigDriver{
		AppendPVC:                    make(map[string]bool),
		PersistentVolumeClaimDrivers: make(map[string]v1Core.PersistentVolumeClaim),
	}
	for _, pvcTemplate := range sts.PVCTemplates {
		storageReq := v1Core.ResourceList{"storage": resource.MustParse(pvcTemplate.StorageSizeRequest)}
		pvc := v1Core.PersistentVolumeClaim{
			Spec: v1Core.PersistentVolumeClaimSpec{
				AccessModes: []v1Core.PersistentVolumeAccessMode{v1Core.PersistentVolumeAccessMode(pvcTemplate.AccessMode)},
				Resources: v1Core.ResourceRequirements{
					Requests: storageReq,
				},
				VolumeName: pvcTemplate.Name,
			},
		}
		pvcCfg.AppendPVC[pvcTemplate.Name] = true
		pvcCfg.PersistentVolumeClaimDrivers[pvcTemplate.Name] = pvc
	}
	stsDriver.PVCDriver = &pvcCfg
	return stsDriver, nil
}

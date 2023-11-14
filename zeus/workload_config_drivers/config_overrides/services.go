package config_overrides

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ServiceDriver struct {
	v1.Service
	ExtendPorts []v1.ServicePort
}

func (s *ServiceDriver) SetServiceConfigs(svc *v1.Service) {
	if svc == nil {
		return
	}
	if s.Service.Spec.Ports != nil {
		svc.Spec.Ports = s.Service.Spec.Ports
	}
	if s.ExtendPorts != nil {
		if svc.Spec.Ports == nil {
			svc.Spec.Ports = []v1.ServicePort{}
		}
		svc.Spec.Ports = append(svc.Spec.Ports, s.ExtendPorts...)
	}
}

func (s *ServiceDriver) AddNginxTargetPort(portName, targetPortName string) {
	if s.ExtendPorts == nil {
		s.ExtendPorts = []v1.ServicePort{}
	}
	s.ExtendPorts = append(s.ExtendPorts, v1.ServicePort{
		Name:       portName,
		Protocol:   "TCP",
		Port:       80,
		TargetPort: intstr.IntOrString{Type: intstr.String, StrVal: targetPortName},
	})
}

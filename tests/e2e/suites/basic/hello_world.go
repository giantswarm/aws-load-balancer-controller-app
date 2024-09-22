package basic

import (
	"slices"

	"github.com/giantswarm/apptest-framework/pkg/state"
	"github.com/giantswarm/clustertest/pkg/logger"
	v1 "k8s.io/api/core/v1"
	v2 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/types"
)

func serviceHasLBHostnameSetInStatus(wcClient client.Client, serviceNamespace, serviceName string) (bool, error) {
	logger.Log("Checking if Service has load balancer set in status")
	service := v1.Service{}
	err := wcClient.Get(state.GetContext(), types.NamespacedName{Name: serviceName, Namespace: serviceNamespace}, &service)
	if err != nil {
		logger.Log("Failed to get Service: %v", err)
		return false, err
	}

	if service.Status.LoadBalancer.Ingress != nil &&
		len(service.Status.LoadBalancer.Ingress) > 0 &&
		service.Status.LoadBalancer.Ingress[0].Hostname != "" {

		logger.Log("Load balancer hostname found in Service status: %s", service.Status.LoadBalancer.Ingress[0].Hostname)
		return true, nil
	}

	return false, nil
}

func createServiceLoadBalancer(wcClient client.Client, serviceNamespace, serviceName string) (*v1.Service, error) {
	logger.Log("Creating a Service of type LoadBalancer")
	service := &v1.Service{
		ObjectMeta: v2.ObjectMeta{
			Name:      serviceName,
			Namespace: serviceNamespace,
		},
		Spec: v1.ServiceSpec{
			Type: v1.ServiceTypeLoadBalancer,
			Ports: []v1.ServicePort{
				{
					Port: 54321,
				},
			},
			Selector: map[string]string{
				"app": "not-used",
			},
		},
	}

	err := wcClient.Create(state.GetContext(), service)
	if err != nil {
		logger.Log("Failed to create Service: %v", err)
		return &v1.Service{}, err
	}

	return service, nil
}

func serviceHasFinalizer(wcClient client.Client, serviceNamespace, serviceName, finalizer string) (bool, error) {
	logger.Log("Checking if Service has the aws-load-balancer-controller finalizer")
	service := &v1.Service{}
	err := wcClient.Get(state.GetContext(), types.NamespacedName{Name: serviceName, Namespace: serviceNamespace}, service)
	if err != nil {
		logger.Log("Failed to get Service: %v", err)
		return false, err
	}

	return slices.Contains(service.GetFinalizers(), finalizer), nil
}

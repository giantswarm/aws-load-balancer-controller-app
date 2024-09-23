package basic

import (
	"slices"
	"testing"
	"time"

	"github.com/giantswarm/apptest-framework/pkg/state"
	"github.com/giantswarm/clustertest/pkg/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/apptest-framework/pkg/config"
	"github.com/giantswarm/apptest-framework/pkg/suite"
)

const (
	isUpgrade = false
)

func TestBasic(t *testing.T) {
	suite.New(config.MustLoad("../../config.yaml")).
		// The namespace to install the app into within the workload cluster
		WithInstallNamespace("kube-system").
		// If this is an upgrade test or not.
		// If true, the suite will first install the latest released version of the app before upgrading to the test version
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		Tests(func() {
			var service *corev1.Service
			It("should manage LB creation", func() {
				wcClient, err := state.GetFramework().WC(state.GetCluster().Name)
				Expect(err).Should(Succeed())

				Eventually(func() error {
					service, err = createServiceLoadBalancer(wcClient, "default", "test-aws-lb-controller")
					return err
				}).
					WithTimeout(2 * time.Minute).
					WithPolling(5 * time.Second).
					Should(Succeed())

				Eventually(func() (bool, error) {
					return serviceHasLBHostnameSetInStatus(wcClient, service.Namespace, service.Name)
				}).
					WithTimeout(6 * time.Minute).
					WithPolling(5 * time.Second).
					Should(BeTrueBecause("We expect the LoadBalancer hostname to be set in the Service status"))

				// We make sure the `Service` has the finalizer added by the aws-load-balancer-controller
				Eventually(func() (bool, error) {
					return serviceHasFinalizer(wcClient, service.Namespace, service.Name, "service.k8s.aws/resources")
				}).
					WithTimeout(3 * time.Minute).
					WithPolling(5 * time.Second).
					Should(BeTrueBecause("We expect the finalizer to be added to the Service by the aws-load-balancer-controller"))
			})
		}).
		Run(t, "Basic Test")
}

func serviceHasLBHostnameSetInStatus(wcClient client.Client, serviceNamespace, serviceName string) (bool, error) {
	logger.Log("Checking if Service has load balancer set in status")
	service := corev1.Service{}
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

func createServiceLoadBalancer(wcClient client.Client, serviceNamespace, serviceName string) (*corev1.Service, error) {
	logger.Log("Creating a Service of type LoadBalancer")
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: serviceNamespace,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeLoadBalancer,
			Ports: []corev1.ServicePort{
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
		return &corev1.Service{}, err
	}

	return service, nil
}

func serviceHasFinalizer(wcClient client.Client, serviceNamespace, serviceName, finalizer string) (bool, error) {
	logger.Log("Checking if Service has the aws-load-balancer-controller finalizer")
	service := &corev1.Service{}
	err := wcClient.Get(state.GetContext(), types.NamespacedName{Name: serviceName, Namespace: serviceNamespace}, service)
	if err != nil {
		logger.Log("Failed to get Service: %v", err)
		return false, err
	}

	return slices.Contains(service.GetFinalizers(), finalizer), nil
}

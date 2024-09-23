package basic

import (
	"testing"
	"time"

	"github.com/giantswarm/apptest-framework/pkg/state"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

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
			var service *v1.Service
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

package basic

import (
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/giantswarm/clustertest/pkg/logger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/giantswarm/apptest-framework/pkg/config"
	"github.com/giantswarm/apptest-framework/pkg/state"
	"github.com/giantswarm/apptest-framework/pkg/suite"
	"github.com/giantswarm/clustertest/pkg/wait"
)

const (
	isUpgrade        = false
	appReadyTimeout  = 5 * time.Minute
	appReadyInterval = 5 * time.Second
)

func TestBasic(t *testing.T) {
	suite.New(config.MustLoad("../../config.yaml")).
		// The namespace to install the app into within the workload cluster
		WithInstallNamespace("kube-system").
		// If this is an upgrade test or not.
		// If true, the suite will first install the latest released version of the app before upgrading to the test version
		WithIsUpgrade(isUpgrade).
		WithValuesFile("./values.yaml").
		AfterClusterReady(func() {
		}).
		BeforeUpgrade(func() {
			// Perform any checks between installing the latest released version
			// and upgrading it to the version to test
			// E.g. ensure that the initial install has completed and has settled before upgrading
		}).
		Tests(func() {
			var (
				helloWorldIngressHost string
				helloWorldIngressUrl  string
			)
			BeforeEach(func() {
				helloWorldIngressHost = fmt.Sprintf("hello-world.%s", getWorkloadClusterBaseDomain())
				helloWorldIngressUrl = fmt.Sprintf("https://%s", helloWorldIngressHost)
			})

			It("should serve traffic from hello-world", func() {
				By("creating the ingress-nginx app CR")

				nginxIngressApp, err := newNginxIngressApp()
				Expect(err).To(BeNil())

				Eventually(wait.IsAppDeployed(state.GetContext(), state.GetFramework().MC(), nginxIngressApp.InstallName, nginxIngressApp.GetNamespace())).
					WithTimeout(appReadyTimeout).
					WithPolling(appReadyInterval).
					Should(BeTrue())

				By("creating the hello-world app CR")

				helloWorldApp, err := newHelloWorldApp(helloWorldIngressHost)
				Expect(err).To(BeNil())
				Eventually(func() (bool, error) {
					// The `ingress-nginx` app has an admission webhooks for `Ingress` resources. While the controller
					// is starting, the webhooks are not available, causing any Ingress operation to fail.
					// The `hello-world-app` fails to install during this time because it creates an Ingress resource.
					// We must wait until the deployment succeeds.

					return patchAndWait(helloWorldApp)
				}).
					WithTimeout(6 * time.Minute).
					WithPolling(5 * time.Second).
					Should(BeTrue())

				By("adding a LB to the ingress CR")

				Eventually(func() (bool, error) {
					return ingressHasLB("default", "hello-world")
				}).
					WithTimeout(6 * time.Minute).
					WithPolling(5 * time.Second).
					Should(BeTrue())

				By("adding aws-load-balancer-controller finalizer to service")

				// We make sure the `Service` has the finalizer added by the aws-load-balancer-controller
				Eventually(func() (bool, error) {
					return serviceHasFinalizer(nginxIngressApp.GetInstallNamespace(), "ingress-nginx-controller", "service.k8s.aws/resources")
				}).
					WithTimeout(6 * time.Minute).
					WithPolling(5 * time.Second).
					Should(BeTrue())

				By("serving responses from the backend")

				httpClient := newHttpClientWithProxy()
				Eventually(func() (string, error) {
					logger.Log("Trying to get a successful response from %s", helloWorldIngressUrl)
					resp, err := httpClient.Get(helloWorldIngressUrl)
					if err != nil {
						return "", err
					}
					defer resp.Body.Close()

					if resp.StatusCode != http.StatusOK {
						logger.Log("Was expecting status code '%d' but actually got '%d'", http.StatusOK, resp.StatusCode)
						return "", err
					}

					bodyBytes, err := io.ReadAll(resp.Body)
					if err != nil {
						logger.Log("Was not expecting the response body to be empty")
						return "", err
					}

					return string(bodyBytes), nil
				}).
					WithTimeout(15 * time.Minute).
					WithPolling(5 * time.Second).
					Should(ContainSubstring("Hello World"))

				By("uninstalling the hello-world app")

				err = state.GetFramework().MC().DeleteApp(state.GetContext(), *helloWorldApp)
				Expect(err).ShouldNot(HaveOccurred())
			})
		}).
		Run(t, "Basic Test")
}

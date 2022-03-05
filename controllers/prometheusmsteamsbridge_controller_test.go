package controllers

import (
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shibumi/prometheus-msteams-operator/api/v1alpha1"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
)

var _ = Describe("Prometheus MSTeams Bridge Controller", func() {
	const (
		Name      = "test-bridge"
		Namespace = "default"
		Replicas  = 1
		Image     = "quay.io/prometheusmsteams/prometheus-msteams:v1.5.0"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)
	Context("When Prometheus MSTeams Bridge CR gets created", func() {
		It("should create the deployment and service", func() {
			By("creating the necessary CR")
			ctx := context.Background()
			bridge := &v1alpha1.PrometheusMSTeamsBridge{
				ObjectMeta: metav1.ObjectMeta{
					Name:      Name,
					Namespace: Namespace,
				},
				Spec: v1alpha1.PrometheusMSTeamsBridgeSpec{
					Replicas: Replicas,
					Image:    Image,
				},
				Status: v1alpha1.PrometheusMSTeamsBridgeStatus{},
			}
			Expect(k8sClient.Create(ctx, bridge)).Should(Succeed())

			crLookup := types.NamespacedName{
				Name:      Name,
				Namespace: Namespace,
			}
			createdCR := &v1alpha1.PrometheusMSTeamsBridge{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, crLookup, createdCR)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())

			By("checking if the deployment exists")
			deployLookup := types.NamespacedName{
				Namespace: Namespace,
				Name:      "bridge-pod",
			}
			createdDeploy := &v1.Deployment{}
			Eventually(func() bool {
				err := k8sClient.Get(ctx, deployLookup, createdDeploy)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())

			By("checking if the service got created")
			serviceLookup := types.NamespacedName{
				Namespace: Namespace,
				Name:      "bridge-service",
			}
			createdService := &corev1.Service{}

			Eventually(func() bool {
				err := k8sClient.Get(ctx, serviceLookup, createdService)
				if err != nil {
					return false
				}
				return true
			}, timeout, interval).Should(BeTrue())
		})
	})
})

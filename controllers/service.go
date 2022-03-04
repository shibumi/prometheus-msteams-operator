package controllers

import (
	"context"

	bridgev1alpha1 "github.com/shibumi/prometheus-msteams-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ensureService ensures that a service has been created for our CR
func (r *PrometheusMSTeamsBridgeReconciler) ensureService(request reconcile.Request, instance *bridgev1alpha1.PrometheusMSTeamsBridge, service *corev1.Service) (*reconcile.Result, error) {

	found := &corev1.Service{}
	err := r.Get(context.Background(), types.NamespacedName{
		Name:      service.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {
		err = r.Create(context.Background(), service)

		if err != nil {
			return &reconcile.Result{}, err
		}
		return nil, nil
	}
	if err != nil {
		return &reconcile.Result{}, err
	}

	return nil, nil
}

// createService creates the necessary service for the deployment
func (r *PrometheusMSTeamsBridgeReconciler) createService(v *bridgev1alpha1.PrometheusMSTeamsBridge) *corev1.Service {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bridge-service",
			Namespace: v.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "bridge"},
			Ports: []corev1.ServicePort{{
				Protocol:   corev1.ProtocolTCP,
				Port:       80,
				TargetPort: intstr.FromInt(80),
			}},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	controllerutil.SetControllerReference(v, service, r.Scheme)
	return service
}

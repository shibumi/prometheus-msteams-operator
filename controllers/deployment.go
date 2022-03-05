package controllers

import (
	"context"

	bridgev1alpha1 "github.com/shibumi/prometheus-msteams-operator/api/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ensureDeployment ensures a deployments presence for a CR
func (r *PrometheusMSTeamsBridgeReconciler) ensureDeployment(request reconcile.Request, instance *bridgev1alpha1.PrometheusMSTeamsBridge, deploy *appsv1.Deployment) (*reconcile.Result, error) {
	found := &appsv1.Deployment{}
	err := r.Get(context.Background(), types.NamespacedName{
		Name:      deploy.Name,
		Namespace: instance.Namespace,
	}, found)
	if err != nil && errors.IsNotFound(err) {
		//if deployment could not get found, create the deployment on the cluster
		err = r.Create(context.Background(), deploy)
		if err != nil {
			return &reconcile.Result{}, nil
		}
		// deployment was successful
		return nil, nil
	}
	if err != nil {
		return &reconcile.Result{}, err
	}
	return nil, nil
}

// createDeployment creates the necessary deployment for a CR
func (r *PrometheusMSTeamsBridgeReconciler) createDeployment(b *bridgev1alpha1.PrometheusMSTeamsBridge) *appsv1.Deployment {
	deploy := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bridge-pod",
			Namespace: b.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &b.Spec.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": "bridge"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "bridge"},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:           b.Spec.Image,
						ImagePullPolicy: corev1.PullIfNotPresent,
						Name:            "bridge-pod",
						Ports: []corev1.ContainerPort{{
							ContainerPort: 80,
							Name:          "web",
						}},
					}},
				},
			},
		},
	}

	controllerutil.SetControllerReference(b, deploy, r.Scheme)
	return deploy
}

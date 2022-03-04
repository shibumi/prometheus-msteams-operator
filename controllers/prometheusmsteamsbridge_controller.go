/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	bridgev1alpha1 "github.com/shibumi/prometheus-msteams-operator/api/v1alpha1"
)

// PrometheusMSTeamsBridgeReconciler reconciles a PrometheusMSTeamsBridge object
type PrometheusMSTeamsBridgeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=bridge.shibumi.dev,resources=prometheusmsteamsbridges,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=bridge.shibumi.dev,resources=prometheusmsteamsbridges/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=bridge.shibumi.dev,resources=prometheusmsteamsbridges/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PrometheusMSTeamsBridge object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *PrometheusMSTeamsBridgeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx).WithValues("Bridge", req.Namespace)

	instance := &bridgev1alpha1.PrometheusMSTeamsBridge{}
	err := r.Get(context.Background(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	found := &appsv1.Deployment{}
	err = r.Get(context.Background(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, found)
	var result *reconcile.Result
	result, err = r.ensureDeployment(req, instance, r.createDeployment(instance))
	if result != nil {
		log.Error(err, "Deployment not ready")
		return *result, err
	}

	result, err = r.ensureService(req, instance, r.createService(instance))
	if result != nil {
		log.Error(err, "Service not ready")
		return *result, err
	}

	log.Info("Skip reconcile: Deployment already exist",
		"Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PrometheusMSTeamsBridgeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&bridgev1alpha1.PrometheusMSTeamsBridge{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}

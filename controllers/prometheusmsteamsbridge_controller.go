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

// Reconcile waits for a PrometheusMSTeamsBridge CR and creates the necessary resources for this CR
func (r *PrometheusMSTeamsBridgeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx).WithValues("Bridge", req.Namespace)

	// lookup deployed PrometheusMSTeamsBridge CR
	instance := &bridgev1alpha1.PrometheusMSTeamsBridge{}
	err := r.Get(context.Background(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	found := &appsv1.Deployment{}
	// try to look up the deployment
	err = r.Get(context.Background(), types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}, found)
	var result *reconcile.Result
	// ensure Deployment exists.. if not create it
	result, err = r.ensureDeployment(req, instance, r.createDeployment(instance))
	if result != nil {
		l.Error(err, "Deployment not ready")
		return *result, err
	}

	result, err = r.ensureService(req, instance, r.createService(instance))
	if result != nil {
		l.Error(err, "Service not ready")
		return *result, err
	}

	l.Info("Skip reconcile: Deployment already exist",
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

/*
Copyright 2023.

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
	"fmt"
	v1alpha1 "github.com/Technion-SpotOS/EvacuationPlanner/pkg/api/v1alpha1"
	"github.com/go-logr/logr"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const loggerName = "evacuation-planner-controller"

// EvacuationPlannerReconciler reconciles a SpotInstance object.
type EvacuationPlannerReconciler struct {
	log logr.Logger
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *EvacuationPlannerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	return ctrl.Result{}, nil
}

// setupEvacuationPlannerController sets up the controller with the Manager.
func setupEvacuationPlannerController(mgr ctrl.Manager) error {
	if err := ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.EvacuationPlan{}).
		Complete(&EvacuationPlannerReconciler{ctrl.Log.WithName(loggerName), mgr.GetClient(), mgr.GetScheme()}); err != nil {
		return fmt.Errorf("failed to add spot-instance controller to the manager: %w", err)
	}

	return nil
}

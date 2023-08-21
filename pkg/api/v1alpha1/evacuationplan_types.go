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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EvacuationPlanSpec defines the desired state of EvacuationPlan.
type EvacuationPlanSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// SequentialJuggle is a map of workload-name -> instance-name to be evacuated sequentially in appearing order.
	SequentialJuggle string `json:"region,omitempty"`
}

// EvacuationPlanPhase represents the phase of the evacuation plan.
type EvacuationPlanPhase string

const (
	// EvacuationPlanPhasePending is the phase when the plan is created and not yet processed.
	EvacuationPlanPhasePending EvacuationPlanPhase = "Pending"
	// EvacuationPlanPhaseInProgress is the phase when the plan is being processed.
	EvacuationPlanPhaseInProgress EvacuationPlanPhase = "InProgress"
	// EvacuationPlanPhaseCompleted is the phase when the plan is completed.
	EvacuationPlanPhaseCompleted EvacuationPlanPhase = "Completed"
)

// EvacuationPlanStatus defines the observed state of EvacuationPlan.
type EvacuationPlanStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Phase is the current phase of the evacuation plan.
	Phase EvacuationPlanPhase `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// EvacuationPlan is the Schema for the evacuationplan API.
type EvacuationPlan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EvacuationPlanSpec   `json:"spec,omitempty"`
	Status EvacuationPlanStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// EvacuationPlanList contains a list of EvacuationPlan.
type EvacuationPlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EvacuationPlan `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EvacuationPlan{}, &EvacuationPlanList{})
}

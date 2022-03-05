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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrometheusMSTeamsBridgeSpec defines the desired state of PrometheusMSTeamsBridge
type PrometheusMSTeamsBridgeSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
}

// PrometheusMSTeamsBridgeStatus defines the observed state of PrometheusMSTeamsBridge
//+kubebuilder:subresource:status
type PrometheusMSTeamsBridgeStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PrometheusMSTeamsBridge is the Schema for the prometheusmsteamsbridges API
type PrometheusMSTeamsBridge struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrometheusMSTeamsBridgeSpec   `json:"spec,omitempty"`
	Status PrometheusMSTeamsBridgeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PrometheusMSTeamsBridgeList contains a list of PrometheusMSTeamsBridge
type PrometheusMSTeamsBridgeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrometheusMSTeamsBridge `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrometheusMSTeamsBridge{}, &PrometheusMSTeamsBridgeList{})
}

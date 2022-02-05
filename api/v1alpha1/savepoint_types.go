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
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SavepointSpec defines the desired state of Savepoint
type SavepointSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Savepoint. Edit savepoint_types.go to remove/update
	Metadata appmanager_apis.SavepointMetadata `json:"metadata,omitempty"`
	Spec     appmanager_apis.SavepointSpec     `json:"spec,omitempty"`
}

// SavepointStatus defines the observed state of Savepoint
type SavepointStatus struct {
	LastSync metav1.Time `json:"lastSync,omitempty"`
	State    string      `json:"status,omitempty"`
}

// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="LAST SYNC",type="string",JSONPath=".status.lastSync"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.state"

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

type Savepoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SavepointSpec   `json:"spec,omitempty"`
	Status SavepointStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
// SavepointList contains a list of Savepoint
type SavepointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Savepoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Savepoint{}, &SavepointList{})
}

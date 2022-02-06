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
	platform_apis "efrat19.io/vvp-gitops-operator/pkg/platform_apis"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FormatsSpec defines the desired state of Formats
type FormatsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Formats. Edit formats_types.go to remove/update
	Spec     platform_apis.Format     `json:"spec,omitempty"`
}

// FormatsStatus defines the observed state of Formats
type FormatsStatus struct {
	LastSync metav1.Time `json:"lastSync,omitempty"`
	State    string      `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="LAST SYNC",type="date",JSONPath=".status.lastSync"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.state"
// Formats is the Schema for the formats API
type Formats struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FormatsSpec   `json:"spec,omitempty"`
	Status FormatsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FormatsList contains a list of Formats
type FormatsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Formats `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Formats{}, &FormatsList{})
}

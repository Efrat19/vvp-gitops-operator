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

// SqlScriptSpec defines the desired state of SqlScript
type SqlScriptSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Spec     platform_apis.SqlScript     `json:"spec,omitempty"`
}

// SqlScriptsStatus defines the observed state of SqlScripts
type SqlScriptStatus struct {
	LastSync metav1.Time `json:"lastSync,omitempty"`
	State    string      `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="LAST SYNC",type="date",JSONPath=".status.lastSync"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.state"

// SqlScript is the Schema for the sqlscripts API
type SqlScript struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SqlScriptSpec   `json:"spec,omitempty"`
	Status SqlScriptStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SqlScriptList contains a list of SqlScript
type SqlScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlScript `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SqlScript{}, &SqlScriptList{})
}

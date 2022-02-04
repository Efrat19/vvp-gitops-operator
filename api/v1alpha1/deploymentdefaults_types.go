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
	appmanager_apis "efrat19.io/vvp-gitops-operator/pkg/appmanager_apis"

)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeploymentDefaultsSpec defines the desired state of DeploymentDefaults
type DeploymentDefaultsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Metadata appmanager_apis.DeploymentDefaultsMetadata `json:"metadata,omitempty"`
	Spec     appmanager_apis.DeploymentSpec     `json:"spec,omitempty"`
}

// DeploymentDefaultsStatus defines the observed state of DeploymentDefaults
type DeploymentDefaultsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	State   string     
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state"

// DeploymentDefaults is the Schema for the deploymentdefaults API
type DeploymentDefaults struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentDefaultsSpec   `json:"spec,omitempty"`
	Status DeploymentDefaultsStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeploymentDefaultsList contains a list of DeploymentDefaults
type DeploymentDefaultsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentDefaults `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeploymentDefaults{}, &DeploymentDefaultsList{})
}

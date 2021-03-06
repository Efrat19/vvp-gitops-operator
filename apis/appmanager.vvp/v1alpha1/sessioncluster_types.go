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

// SessionClusterSpec defines the desired state of SessionCluster
type SessionClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SessionCluster. Edit sessioncluster_types.go to remove/update
	Metadata appmanager_apis.SessionClusterMetadata `json:"metadata,omitempty"`
	Spec     appmanager_apis.SessionClusterSpec     `json:"spec,omitempty"`
}

// SessionClusterStatus defines the observed state of SessionCluster
type SessionClusterStatus struct {
	LastSync metav1.Time `json:"lastSync,omitempty"`
	State    string      `json:"state,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="LAST SYNC",type="date",JSONPath=".status.lastSync"
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.state"

type SessionCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SessionClusterSpec   `json:"spec,omitempty"`
	Status SessionClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SessionClusterList contains a list of SessionCluster
type SessionClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SessionCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SessionCluster{}, &SessionClusterList{})
}

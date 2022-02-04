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

// JobSpec defines the desired state of Job
type JobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Job. Edit job_types.go to remove/update
	Metadata appmanager_apis.JobMetadata `json:"metadata,omitempty"`
	Spec     appmanager_apis.JobSpec     `json:"spec,omitempty"`
	Status   appmanager_apis.JobStatus   `json:"status,omitempty"`

}

// JobStatus defines the observed state of Job
type JobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Running *appmanager_apis.DeploymentStatusRunning `json:"running,omitempty"`
	Failure      *appmanager_apis.Failure          `json:"failure,omitempty"`
	SinkTables   *[]appmanager_apis.JobTable        `json:"sinkTables,omitempty"`
	SourceTables *[]appmanager_apis.JobTable        `json:"sourceTables,omitempty"`
	Started      *appmanager_apis.JobStatusStarted `json:"started,omitempty"`
	State        string            `json:"state,omitempty"`}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status


// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state"

// Job is the Schema for the jobs API
type Job struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JobSpec   `json:"spec,omitempty"`
	Status JobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// JobList contains a list of Job
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Job `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Job{}, &JobList{})
}

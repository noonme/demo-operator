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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MyDaemonSetSpec defines the desired state of MyDaemonSet
type MyDaemonSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of MyDaemonSet. Edit mydaemonset_types.go to remove/update
	Image string `json:"image,omitempty"`
}

// MyDaemonSetStatus defines the observed state of MyDaemonSet
type MyDaemonSetStatus struct {
	AvaiableReplicas int `json:"avaiableReplicas,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MyDaemonSet is the Schema for the mydaemonsets API
type MyDaemonSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyDaemonSetSpec   `json:"spec,omitempty"`
	Status MyDaemonSetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MyDaemonSetList contains a list of MyDaemonSet
type MyDaemonSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyDaemonSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MyDaemonSet{}, &MyDaemonSetList{})
}

/*
Copyright 2025 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// IAMUserListParameters are the configurable fields of a IAMUserList.
type IAMUserListParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// IAMUserListObservation are the observable fields of a IAMUserList.
type IAMUserListObservation struct {
	ConfigurableField string `json:"configurableField"`
	ObservableField   string `json:"observableField,omitempty"`
}

// A IAMUserListSpec defines the desired state of a IAMUserList.
type IAMUserListSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IAMUserListParameters `json:"forProvider"`
}

// A IAMUserListStatus represents the observed state of a IAMUserList.
type IAMUserListStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IAMUserListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A IAMUserList is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsiamlist}
type IAMUserList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IAMUserListSpec   `json:"spec"`
	Status IAMUserListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IAMUserListList contains a list of IAMUserList
type IAMUserListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAMUserList `json:"items"`
}

// IAMUserList type metadata.
var (
	IAMUserListKind             = reflect.TypeOf(IAMUserList{}).Name()
	IAMUserListGroupKind        = schema.GroupKind{Group: Group, Kind: IAMUserListKind}.String()
	IAMUserListKindAPIVersion   = IAMUserListKind + "." + SchemeGroupVersion.String()
	IAMUserListGroupVersionKind = SchemeGroupVersion.WithKind(IAMUserListKind)
)

func init() {
	SchemeBuilder.Register(&IAMUserList{}, &IAMUserListList{})
}

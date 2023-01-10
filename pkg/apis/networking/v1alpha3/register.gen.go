// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1alpha3

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// Package-wide variables from generator "register".
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha3"}
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

const (
	// Package-wide consts from generator "register".
	GroupName = "networking.istio.io"
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&DestinationRule{},
		&DestinationRuleList{},
		&EnvoyFilter{},
		&EnvoyFilterList{},
		&Gateway{},
		&GatewayList{},
		&OCSPStaple{},
		&OCSPStapleList{},
		&ServiceEntry{},
		&ServiceEntryList{},
		&Sidecar{},
		&SidecarList{},
		&VirtualService{},
		&VirtualServiceList{},
		&WorkloadEntry{},
		&WorkloadEntryList{},
		&WorkloadGroup{},
		&WorkloadGroupList{},
	)
	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

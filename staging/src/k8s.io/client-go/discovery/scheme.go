/*
Copyright 2026 The Kubernetes Authors.

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

package discovery

import (
	apidiscoveryv2 "k8s.io/api/apidiscovery/v2"
	apidiscoveryv2beta1 "k8s.io/api/apidiscovery/v2beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// discoveryScheme is a minimal scheme containing only the types needed by
// the discovery client, avoiding a dependency on the full kubernetes scheme
// which registers all 52+ API group versions.
var discoveryScheme = runtime.NewScheme()

// Codecs provides encoding and decoding for the discovery scheme.
var Codecs = serializer.NewCodecFactory(discoveryScheme)

func init() {
	metav1.AddToGroupVersion(discoveryScheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(apidiscoveryv2.AddToScheme(discoveryScheme))
	utilruntime.Must(apidiscoveryv2beta1.AddToScheme(discoveryScheme))
}

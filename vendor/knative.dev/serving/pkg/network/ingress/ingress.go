/*
Copyright 2019 The Knative Authors.

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

package ingress

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"knative.dev/pkg/network"
	"knative.dev/serving/pkg/apis/networking/v1alpha1"
)

// ComputeHash computes a hash of the Ingress Spec, Namespace and Name
func ComputeHash(ing *v1alpha1.Ingress) ([sha256.Size]byte, error) {
	bytes, err := json.Marshal(ing.Spec)
	if err != nil {
		return [sha256.Size]byte{}, fmt.Errorf("failed to serialize Ingress: %w", err)
	}
	bytes = append(bytes, []byte(ing.GetNamespace())...)
	bytes = append(bytes, []byte(ing.GetName())...)
	return sha256.Sum256(bytes), nil
}

// HostsPerVisibility takes an Ingress and a map from visibility levels to a set of string keys,
// it then returns a map from that key space to the hosts under that visibility.
func HostsPerVisibility(ing *v1alpha1.Ingress, visibilityToKey map[v1alpha1.IngressVisibility]sets.String) map[string]sets.String {
	output := make(map[string]sets.String)
	for _, rule := range ing.Spec.Rules {
		for host := range ExpandedHosts(sets.NewString(rule.Hosts...)) {
			for key := range visibilityToKey[rule.Visibility] {
				if _, ok := output[key]; !ok {
					output[key] = sets.NewString()
				}
				output[key].Insert(host)
			}
		}
	}
	return output
}

// ExpandedHosts sets up hosts for the short-names for cluster DNS names.
func ExpandedHosts(hosts sets.String) sets.String {
	expanded := sets.NewString()
	allowedSuffixes := []string{
		"",
		"." + network.GetClusterDomainName(),
		".svc." + network.GetClusterDomainName(),
	}
	for _, h := range hosts.List() {
		for _, suffix := range allowedSuffixes {
			if strings.HasSuffix(h, suffix) {
				expanded.Insert(strings.TrimSuffix(h, suffix))
			}
		}
	}
	return expanded
}

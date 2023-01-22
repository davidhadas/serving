/*
Copyright 2018 The Knative Authors

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

package resources

import (
	"encoding/json"
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
	appsv1 "k8s.io/api/apps/v1"
)

func overlayDeployment(deployment *appsv1.Deployment, patch jsonpatch.Patch) (*appsv1.Deployment, error) {
	original, err := json.Marshal(deployment)
	if err != nil {
		return nil, fmt.Errorf("marshaling deployment: %w", err)
	}

	modified, err := patch.Apply(original)
	if err != nil {
		return nil, fmt.Errorf("applying patch to deployment: %w", err)
	}

	newDeployment := &appsv1.Deployment{}
	err = json.Unmarshal(modified, newDeployment)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling patched deployment: %w", err)
	}
	return newDeployment, nil
}

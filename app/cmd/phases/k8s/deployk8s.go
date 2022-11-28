/*
Copyright 2019 The Kubernetes Authors.

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

package phases

import (
	"fmt"

	"github.com/283713406/kins/app/cmd/phases/workflow"
	"github.com/283713406/kins/app/phases/deploy"
)

// NewDeployK8sPhase creates a kins workflow phase for deploy-k8s
func NewDeployK8sPhase() workflow.Phase {
	return workflow.Phase{
		Name:  "deploy-k8s",
		Short: "Deploy k8s.",
		Long:  "Deploy kylin container cloud platform.",
		Run:   runDeployK8sPhase,
	}
}

// runDeployK8sPhase executes deploy logic.
func runDeployK8sPhase(c workflow.RunData) error {
	fmt.Println("[mount] Running load k8s images")
	if err := deploy.DeployK8s(); err != nil {
		return err
	}
	return nil
}
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
	"github.com/283713406/kins/app/phases/mounts"
)

// NewMountK8sIsoPhase creates a kins workflow phase for mount-k8s-iso
func NewMountK8sIsoPhase() workflow.Phase {
	return workflow.Phase{
		Name:  "mount-k8s-iso",
		Short: "Mount k8s iso to local.",
		Long:  "Mount k8s images file to local path.",
		Run:   runMountK8sIsoPhase,
	}
}

// runMountK8sIsoPhase executes mount logic.
func runMountK8sIsoPhase(c workflow.RunData) error {
	fmt.Println("[mount] Running mount k8s iso")
	if err := mounts.RunMountK8sIso(); err != nil {
		return err
	}
	return nil
}
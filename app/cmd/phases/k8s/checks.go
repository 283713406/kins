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
	"github.com/283713406/kins/app/utils/normalizer"

	"github.com/283713406/kins/app/cmd/phases/workflow"
	"github.com/283713406/kins/app/phases/checks"
)

var (
	checkExample = normalizer.Examples(`
		# Run pre-flight checks for kins k8s.
		kubeadm init phase checks
		`)
)

// NewChecksPhase creates a kins workflow phase implements pre checks for install
func NewChecksPhase() workflow.Phase {
	return workflow.Phase{
		Name:    "check gateway",
		Aliases: []string{"check-gw"},
		Short:   "Run check gateway",
		Long:    "Run check-gw check for gateway.",
		Example: checkExample,
		Run:     runChecks,
	}
}

// runChecks executes checks logic.
func runChecks(c workflow.RunData) error {
	fmt.Println("[checks] Running checks")
	if err := checks.RunPreChecks(); err != nil {
		return err
	}
	return nil
}

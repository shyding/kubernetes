// +build !windows

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

package dockershim

import (
	dockertypes "github.com/docker/docker/api/types"
	runtimeapi "k8s.io/kubernetes/pkg/kubelet/apis/cri/runtime/v1alpha2"
)

type containerCreationCleanupInfo struct{}

// applyPlatformSpecificDockerConfig applies platform-specific configurations to a dockertypes.ContainerCreateConfig struct.
// The containerCreationCleanupInfo struct it returns will be passed as is to performPlatformSpecificContainerCreationCleanup
// after the container has been created.
func (ds *dockerService) applyPlatformSpecificDockerConfig(*runtimeapi.CreateContainerRequest, *dockertypes.ContainerCreateConfig) (*containerCreationCleanupInfo, error) {
	return nil, nil
}

// performPlatformSpecificContainerCreationCleanup is responsible for doing any platform-specific cleanup
// after a container creation. Any errors it returns are simply logged, but do not fail the container
// creation.
func (ds *dockerService) performPlatformSpecificContainerCreationCleanup(cleanupInfo *containerCreationCleanupInfo) (errors []error) {
	return
}

// platformSpecificContainerCreationInitCleanup is called when dockershim
// is starting, and is meant to clean up any cruft left by previous runs
// creating containers.
// Errors are simply logged, but don't prevent dockershim from starting.
func (ds *dockerService) platformSpecificContainerCreationInitCleanup() (errors []error) {
	return
}

/*
Copyright 2018 Mirantis

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

package proxy

import (
	"google.golang.org/grpc"

	runtimeapi "github.com/nxsre/criproxy/pkg/runtimeapis/v1_15"
)

// ---

type PodSandbox_115 struct {
	inner *runtimeapi.PodSandbox
}

var _ PodSandbox = &PodSandbox_115{}

func (o *PodSandbox_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandbox{}
	} else {
		o.inner = v.(*runtimeapi.PodSandbox)
	}
}
func (o *PodSandbox_115) Unwrap() interface{} { return o.inner }
func (o *PodSandbox_115) Copy() PodSandbox    { r := *o.inner; return &PodSandbox_115{&r} }
func (o *PodSandbox_115) Id() string          { return o.inner.Id }
func (o *PodSandbox_115) SetId(id string)     { o.inner.Id = id }

type Container_115 struct {
	inner *runtimeapi.Container
}

// ---

var _ Container = &Container_115{}

func (o *Container_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Container{}
	} else {
		o.inner = v.(*runtimeapi.Container)
	}
}
func (o *Container_115) Unwrap() interface{}       { return o.inner }
func (o *Container_115) Copy() Container           { r := *o.inner; return &Container_115{&r} }
func (o *Container_115) Id() string                { return o.inner.Id }
func (o *Container_115) SetId(id string)           { o.inner.Id = id }
func (o *Container_115) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *Container_115) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *Container_115) Image() string             { return o.inner.Image.GetImage() }
func (o *Container_115) SetImage(image string)     { o.inner.Image = &runtimeapi.ImageSpec{Image: image} }

// ---

type Image_115 struct {
	inner *runtimeapi.Image
}

var _ Image = &Image_115{}

func (o *Image_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.Image{}
	} else {
		o.inner = v.(*runtimeapi.Image)
	}
}
func (o *Image_115) Unwrap() interface{}                 { return o.inner }
func (o *Image_115) Copy() Image                         { r := *o.inner; return &Image_115{&r} }
func (o *Image_115) Id() string                          { return o.inner.Id }
func (o *Image_115) SetId(id string)                     { o.inner.Id = id }
func (o *Image_115) RepoTags() []string                  { return o.inner.RepoTags }
func (o *Image_115) SetRepoTags(repoTags []string)       { o.inner.RepoTags = repoTags }
func (o *Image_115) RepoDigests() []string               { return o.inner.RepoDigests }
func (o *Image_115) SetRepoDigests(repoDigests []string) { o.inner.RepoDigests = repoDigests }

// ---

type PodSandboxStatus_115 struct {
	inner *runtimeapi.PodSandboxStatus
}

var _ PodSandboxStatus = &PodSandboxStatus_115{}

func (o *PodSandboxStatus_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatus{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatus)
	}
}
func (o *PodSandboxStatus_115) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatus_115) Copy() PodSandboxStatus {
	r := *o.inner
	return &PodSandboxStatus_115{&r}
}
func (o *PodSandboxStatus_115) Id() string      { return o.inner.Id }
func (o *PodSandboxStatus_115) SetId(id string) { o.inner.Id = id }

// ---

type ContainerStatus_115 struct {
	inner *runtimeapi.ContainerStatus
}

var _ ContainerStatus = &ContainerStatus_115{}

func (o *ContainerStatus_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatus{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatus)
	}
}
func (o *ContainerStatus_115) Unwrap() interface{}   { return o.inner }
func (o *ContainerStatus_115) Copy() ContainerStatus { r := *o.inner; return &ContainerStatus_115{&r} }
func (o *ContainerStatus_115) Id() string            { return o.inner.Id }
func (o *ContainerStatus_115) SetId(id string)       { o.inner.Id = id }
func (o *ContainerStatus_115) Image() string         { return o.inner.Image.GetImage() }
func (o *ContainerStatus_115) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ContainerStats_115 struct {
	inner *runtimeapi.ContainerStats
}

var _ ContainerStats = &ContainerStats_115{}

func (o *ContainerStats_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStats{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStats)
	}
}
func (o *ContainerStats_115) Unwrap() interface{}  { return o.inner }
func (o *ContainerStats_115) Copy() ContainerStats { r := *o.inner; return &ContainerStats_115{&r} }
func (o *ContainerStats_115) Id() string           { return o.inner.Attributes.GetId() }
func (o *ContainerStats_115) SetId(id string) {
	if o.inner.Attributes == nil {
		o.inner.Attributes = &runtimeapi.ContainerAttributes{Id: id}
	} else {
		o.inner.Attributes.Id = id
	}
}

// ---

type FilesystemUsage_115 struct {
	inner *runtimeapi.FilesystemUsage
}

func (o *FilesystemUsage_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.FilesystemUsage{}
	} else {
		o.inner = v.(*runtimeapi.FilesystemUsage)
	}
}
func (o *FilesystemUsage_115) Unwrap() interface{} { return o.inner }

// ---

type VersionRequest_115 struct {
	inner *runtimeapi.VersionRequest
}

var _ VersionRequest = &VersionRequest_115{}

func (o *VersionRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionRequest{}
	} else {
		o.inner = v.(*runtimeapi.VersionRequest)
	}
}
func (o *VersionRequest_115) Unwrap() interface{} { return o.inner }

// ---

type VersionResponse_115 struct {
	inner *runtimeapi.VersionResponse
}

var _ VersionResponse = &VersionResponse_115{}

func (o *VersionResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.VersionResponse{}
	} else {
		o.inner = v.(*runtimeapi.VersionResponse)
	}
}
func (o *VersionResponse_115) Unwrap() interface{} { return o.inner }

// ---

type StatusRequest_115 struct {
	inner *runtimeapi.StatusRequest
}

var _ StatusRequest = &StatusRequest_115{}

func (o *StatusRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.StatusRequest)
	}
}
func (o *StatusRequest_115) Unwrap() interface{} { return o.inner }

// ---

type StatusResponse_115 struct {
	inner *runtimeapi.StatusResponse
}

var _ StatusResponse = &StatusResponse_115{}

func (o *StatusResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.StatusResponse)
	}
}
func (o *StatusResponse_115) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigRequest_115 struct {
	inner *runtimeapi.UpdateRuntimeConfigRequest
}

var _ UpdateRuntimeConfigRequest = &UpdateRuntimeConfigRequest_115{}

func (o *UpdateRuntimeConfigRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigRequest)
	}
}
func (o *UpdateRuntimeConfigRequest_115) Unwrap() interface{} { return o.inner }

// ---

type UpdateRuntimeConfigResponse_115 struct {
	inner *runtimeapi.UpdateRuntimeConfigResponse
}

var _ UpdateRuntimeConfigResponse = &UpdateRuntimeConfigResponse_115{}

func (o *UpdateRuntimeConfigResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateRuntimeConfigResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateRuntimeConfigResponse)
	}
}
func (o *UpdateRuntimeConfigResponse_115) Unwrap() interface{} { return o.inner }

// ---

type RunPodSandboxRequest_115 struct {
	inner *runtimeapi.RunPodSandboxRequest
}

var _ RunPodSandboxRequest = &RunPodSandboxRequest_115{}

func (o *RunPodSandboxRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxRequest)
	}
}
func (o *RunPodSandboxRequest_115) Unwrap() interface{} { return o.inner }
func (o *RunPodSandboxRequest_115) GetAnnotations() map[string]string {
	return o.inner.Config.GetAnnotations()
}

// ---

type RunPodSandboxResponse_115 struct {
	inner *runtimeapi.RunPodSandboxResponse
}

var _ RunPodSandboxResponse = &RunPodSandboxResponse_115{}

func (o *RunPodSandboxResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RunPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RunPodSandboxResponse)
	}
}
func (o *RunPodSandboxResponse_115) Unwrap() interface{}       { return o.inner }
func (o *RunPodSandboxResponse_115) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RunPodSandboxResponse_115) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type ListPodSandboxRequest_115 struct {
	inner *runtimeapi.ListPodSandboxRequest
}

var _ ListPodSandboxRequest = &ListPodSandboxRequest_115{}

func (o *ListPodSandboxRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxRequest)
	}
}
func (o *ListPodSandboxRequest_115) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxRequest_115) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListPodSandboxRequest_115) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.PodSandboxFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

// ---

type ListPodSandboxResponse_115 struct {
	inner *runtimeapi.ListPodSandboxResponse
}

var _ ListPodSandboxResponse = &ListPodSandboxResponse_115{}

func (o *ListPodSandboxResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListPodSandboxResponse)
	}
}
func (o *ListPodSandboxResponse_115) Unwrap() interface{} { return o.inner }
func (o *ListPodSandboxResponse_115) Items() []CRIObject {
	var r []CRIObject
	for _, sandbox := range o.inner.Items {
		r = append(r, &PodSandbox_115{sandbox})
	}
	return r
}
func (o *ListPodSandboxResponse_115) SetItems(items []CRIObject) {
	o.inner.Items = nil
	for _, wrapped := range items {
		o.inner.Items = append(o.inner.Items, wrapped.Unwrap().(*runtimeapi.PodSandbox))
	}
}

// ---

type StopPodSandboxRequest_115 struct {
	inner *runtimeapi.StopPodSandboxRequest
}

var _ StopPodSandboxRequest = &StopPodSandboxRequest_115{}

func (o *StopPodSandboxRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxRequest)
	}
}
func (o *StopPodSandboxRequest_115) Unwrap() interface{}       { return o.inner }
func (o *StopPodSandboxRequest_115) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *StopPodSandboxRequest_115) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type StopPodSandboxResponse_115 struct {
	inner *runtimeapi.StopPodSandboxResponse
}

var _ StopPodSandboxResponse = &StopPodSandboxResponse_115{}

func (o *StopPodSandboxResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopPodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopPodSandboxResponse)
	}
}
func (o *StopPodSandboxResponse_115) Unwrap() interface{} { return o.inner }

// ---

type RemovePodSandboxRequest_115 struct {
	inner *runtimeapi.RemovePodSandboxRequest
}

var _ RemovePodSandboxRequest = &RemovePodSandboxRequest_115{}

func (o *RemovePodSandboxRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxRequest)
	}
}
func (o *RemovePodSandboxRequest_115) Unwrap() interface{}       { return o.inner }
func (o *RemovePodSandboxRequest_115) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *RemovePodSandboxRequest_115) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type RemovePodSandboxResponse_115 struct {
	inner *runtimeapi.RemovePodSandboxResponse
}

var _ RemovePodSandboxResponse = &RemovePodSandboxResponse_115{}

func (o *RemovePodSandboxResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemovePodSandboxResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemovePodSandboxResponse)
	}
}
func (o *RemovePodSandboxResponse_115) Unwrap() interface{} { return o.inner }

// ---

type PodSandboxStatusRequest_115 struct {
	inner *runtimeapi.PodSandboxStatusRequest
}

var _ PodSandboxStatusRequest = &PodSandboxStatusRequest_115{}

func (o *PodSandboxStatusRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusRequest)
	}
}
func (o *PodSandboxStatusRequest_115) Unwrap() interface{}       { return o.inner }
func (o *PodSandboxStatusRequest_115) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *PodSandboxStatusRequest_115) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }

// ---

type PodSandboxStatusResponse_115 struct {
	inner *runtimeapi.PodSandboxStatusResponse
}

var _ PodSandboxStatusResponse = &PodSandboxStatusResponse_115{}

func (o *PodSandboxStatusResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PodSandboxStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.PodSandboxStatusResponse)
	}
}
func (o *PodSandboxStatusResponse_115) Unwrap() interface{} { return o.inner }
func (o *PodSandboxStatusResponse_115) Status() PodSandboxStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &PodSandboxStatus_115{o.inner.Status}
}

// ---

type CreateContainerRequest_115 struct {
	inner *runtimeapi.CreateContainerRequest
}

var _ CreateContainerRequest = &CreateContainerRequest_115{}

func (o *CreateContainerRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerRequest)
	}
}
func (o *CreateContainerRequest_115) Unwrap() interface{}       { return o.inner }
func (o *CreateContainerRequest_115) PodSandboxId() string      { return o.inner.PodSandboxId }
func (o *CreateContainerRequest_115) SetPodSandboxId(id string) { o.inner.PodSandboxId = id }
func (o *CreateContainerRequest_115) Image() string {
	if o.inner.Config == nil {
		return ""
	}
	return o.inner.Config.Image.GetImage()
}

func (o *CreateContainerRequest_115) SetImage(image string) {
	if o.inner.Config != nil {
		o.inner.Config.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type CreateContainerResponse_115 struct {
	inner *runtimeapi.CreateContainerResponse
}

var _ CreateContainerResponse = &CreateContainerResponse_115{}

func (o *CreateContainerResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.CreateContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.CreateContainerResponse)
	}
}
func (o *CreateContainerResponse_115) Unwrap() interface{}      { return o.inner }
func (o *CreateContainerResponse_115) ContainerId() string      { return o.inner.ContainerId }
func (o *CreateContainerResponse_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ListContainersRequest_115 struct {
	inner *runtimeapi.ListContainersRequest
}

var _ ListContainersRequest = &ListContainersRequest_115{}

func (o *ListContainersRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersRequest)
	}
}
func (o *ListContainersRequest_115) Unwrap() interface{} { return o.inner }
func (o *ListContainersRequest_115) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainersRequest_115) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainersRequest_115) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainersRequest_115) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainersResponse_115 struct {
	inner *runtimeapi.ListContainersResponse
}

var _ ListContainersResponse = &ListContainersResponse_115{}

func (o *ListContainersResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainersResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainersResponse)
	}
}
func (o *ListContainersResponse_115) Unwrap() interface{} { return o.inner }
func (o *ListContainersResponse_115) Items() []CRIObject {
	var r []CRIObject
	for _, container := range o.inner.Containers {
		r = append(r, &Container_115{container})
	}
	return r
}
func (o *ListContainersResponse_115) SetItems(items []CRIObject) {
	o.inner.Containers = nil
	for _, wrapped := range items {
		o.inner.Containers = append(o.inner.Containers, wrapped.Unwrap().(*runtimeapi.Container))
	}
}

// ---

type ListContainerStatsRequest_115 struct {
	inner *runtimeapi.ListContainerStatsRequest
}

var _ ListContainerStatsRequest = &ListContainerStatsRequest_115{}

func (o *ListContainerStatsRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsRequest)
	}
}
func (o *ListContainerStatsRequest_115) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsRequest_115) IdFilter() string {
	return o.inner.Filter.GetId()
}

func (o *ListContainerStatsRequest_115) SetIdFilter(id string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: id}
	} else {
		o.inner.Filter.Id = id
	}
}

func (o *ListContainerStatsRequest_115) PodSandboxIdFilter() string {
	return o.inner.Filter.GetPodSandboxId()
}

func (o *ListContainerStatsRequest_115) SetPodSandboxIdFilter(podSandboxId string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ContainerStatsFilter{Id: podSandboxId}
	} else {
		o.inner.Filter.PodSandboxId = podSandboxId
	}
}

// ---

type ListContainerStatsResponse_115 struct {
	inner *runtimeapi.ListContainerStatsResponse
}

var _ ListContainerStatsResponse = &ListContainerStatsResponse_115{}

func (o *ListContainerStatsResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListContainerStatsResponse)
	}
}
func (o *ListContainerStatsResponse_115) Unwrap() interface{} { return o.inner }
func (o *ListContainerStatsResponse_115) Items() []CRIObject {
	var r []CRIObject
	for _, stats := range o.inner.Stats {
		r = append(r, &ContainerStats_115{stats})
	}
	return r
}
func (o *ListContainerStatsResponse_115) SetItems(items []CRIObject) {
	o.inner.Stats = nil
	for _, wrapped := range items {
		o.inner.Stats = append(o.inner.Stats, wrapped.Unwrap().(*runtimeapi.ContainerStats))
	}
}

// ---

type StartContainerRequest_115 struct {
	inner *runtimeapi.StartContainerRequest
}

var _ StartContainerRequest = &StartContainerRequest_115{}

func (o *StartContainerRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerRequest)
	}
}
func (o *StartContainerRequest_115) Unwrap() interface{}      { return o.inner }
func (o *StartContainerRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *StartContainerRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StartContainerResponse_115 struct {
	inner *runtimeapi.StartContainerResponse
}

var _ StartContainerResponse = &StartContainerResponse_115{}

func (o *StartContainerResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StartContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StartContainerResponse)
	}
}
func (o *StartContainerResponse_115) Unwrap() interface{} { return o.inner }

// ---

type StopContainerRequest_115 struct {
	inner *runtimeapi.StopContainerRequest
}

var _ StopContainerRequest = &StopContainerRequest_115{}

func (o *StopContainerRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerRequest)
	}
}
func (o *StopContainerRequest_115) Unwrap() interface{}      { return o.inner }
func (o *StopContainerRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *StopContainerRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type StopContainerResponse_115 struct {
	inner *runtimeapi.StopContainerResponse
}

var _ StopContainerResponse = &StopContainerResponse_115{}

func (o *StopContainerResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.StopContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.StopContainerResponse)
	}
}
func (o *StopContainerResponse_115) Unwrap() interface{} { return o.inner }

// ---

type RemoveContainerRequest_115 struct {
	inner *runtimeapi.RemoveContainerRequest
}

var _ RemoveContainerRequest = &RemoveContainerRequest_115{}

func (o *RemoveContainerRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerRequest)
	}
}
func (o *RemoveContainerRequest_115) Unwrap() interface{}      { return o.inner }
func (o *RemoveContainerRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *RemoveContainerRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type RemoveContainerResponse_115 struct {
	inner *runtimeapi.RemoveContainerResponse
}

var _ RemoveContainerResponse = &RemoveContainerResponse_115{}

func (o *RemoveContainerResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveContainerResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveContainerResponse)
	}
}
func (o *RemoveContainerResponse_115) Unwrap() interface{} { return o.inner }

// ---

type ReopenContainerLogRequest_115 struct {
	inner *runtimeapi.ReopenContainerLogRequest
}

var _ ReopenContainerLogRequest = &ReopenContainerLogRequest_115{}

func (o *ReopenContainerLogRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ReopenContainerLogRequest{}
	} else {
		o.inner = v.(*runtimeapi.ReopenContainerLogRequest)
	}
}
func (o *ReopenContainerLogRequest_115) Unwrap() interface{}      { return o.inner }
func (o *ReopenContainerLogRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *ReopenContainerLogRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ReopenContainerLogResponse_115 struct {
	inner *runtimeapi.ReopenContainerLogResponse
}

var _ ReopenContainerLogResponse = &ReopenContainerLogResponse_115{}

func (o *ReopenContainerLogResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ReopenContainerLogResponse{}
	} else {
		o.inner = v.(*runtimeapi.ReopenContainerLogResponse)
	}
}
func (o *ReopenContainerLogResponse_115) Unwrap() interface{} { return o.inner }

// ---

type ContainerStatusRequest_115 struct {
	inner *runtimeapi.ContainerStatusRequest
}

var _ ContainerStatusRequest = &ContainerStatusRequest_115{}

func (o *ContainerStatusRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusRequest)
	}
}
func (o *ContainerStatusRequest_115) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatusRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatusRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatusResponse_115 struct {
	inner *runtimeapi.ContainerStatusResponse
}

var _ ContainerStatusResponse = &ContainerStatusResponse_115{}

func (o *ContainerStatusResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatusResponse)
	}
}
func (o *ContainerStatusResponse_115) Unwrap() interface{} { return o.inner }
func (o *ContainerStatusResponse_115) Status() ContainerStatus {
	if o.inner.Status == nil {
		return nil
	}
	return &ContainerStatus_115{o.inner.Status}
}

// ---

type ContainerStatsRequest_115 struct {
	inner *runtimeapi.ContainerStatsRequest
}

var _ ContainerStatsRequest = &ContainerStatsRequest_115{}

func (o *ContainerStatsRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsRequest{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsRequest)
	}
}
func (o *ContainerStatsRequest_115) Unwrap() interface{}      { return o.inner }
func (o *ContainerStatsRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *ContainerStatsRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ContainerStatsResponse_115 struct {
	inner *runtimeapi.ContainerStatsResponse
}

var _ ContainerStatsResponse = &ContainerStatsResponse_115{}

func (o *ContainerStatsResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ContainerStatsResponse{}
	} else {
		o.inner = v.(*runtimeapi.ContainerStatsResponse)
	}
}
func (o *ContainerStatsResponse_115) Unwrap() interface{} { return o.inner }
func (o *ContainerStatsResponse_115) Stats() ContainerStats {
	if o.inner.Stats == nil {
		return nil
	}
	return &ContainerStats_115{o.inner.Stats}
}

// ---

type ExecSyncRequest_115 struct {
	inner *runtimeapi.ExecSyncRequest
}

var _ ExecSyncRequest = &ExecSyncRequest_115{}

func (o *ExecSyncRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncRequest)
	}
}
func (o *ExecSyncRequest_115) Unwrap() interface{}      { return o.inner }
func (o *ExecSyncRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecSyncRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecSyncResponse_115 struct {
	inner *runtimeapi.ExecSyncResponse
}

var _ ExecSyncResponse = &ExecSyncResponse_115{}

func (o *ExecSyncResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecSyncResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecSyncResponse)
	}
}
func (o *ExecSyncResponse_115) Unwrap() interface{} { return o.inner }

// ---

type ExecRequest_115 struct {
	inner *runtimeapi.ExecRequest
}

var _ ExecRequest = &ExecRequest_115{}

func (o *ExecRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecRequest{}
	} else {
		o.inner = v.(*runtimeapi.ExecRequest)
	}
}
func (o *ExecRequest_115) Unwrap() interface{}      { return o.inner }
func (o *ExecRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *ExecRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type ExecResponse_115 struct {
	inner *runtimeapi.ExecResponse
}

var _ ExecResponse = &ExecResponse_115{}

func (o *ExecResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ExecResponse{}
	} else {
		o.inner = v.(*runtimeapi.ExecResponse)
	}
}
func (o *ExecResponse_115) Unwrap() interface{} { return o.inner }
func (o *ExecResponse_115) Url() string         { return o.inner.Url }
func (o *ExecResponse_115) SetUrl(url string)   { o.inner.Url = url }

// ---

type AttachRequest_115 struct {
	inner *runtimeapi.AttachRequest
}

var _ AttachRequest = &AttachRequest_115{}

func (o *AttachRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachRequest{}
	} else {
		o.inner = v.(*runtimeapi.AttachRequest)
	}
}
func (o *AttachRequest_115) Unwrap() interface{}      { return o.inner }
func (o *AttachRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *AttachRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// ---

type AttachResponse_115 struct {
	inner *runtimeapi.AttachResponse
}

var _ AttachResponse = &AttachResponse_115{}

func (o *AttachResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.AttachResponse{}
	} else {
		o.inner = v.(*runtimeapi.AttachResponse)
	}
}
func (o *AttachResponse_115) Unwrap() interface{} { return o.inner }
func (o *AttachResponse_115) Url() string         { return o.inner.Url }
func (o *AttachResponse_115) SetUrl(url string)   { o.inner.Url = url }

// ---

type PortForwardRequest_115 struct {
	inner *runtimeapi.PortForwardRequest
}

var _ PortForwardRequest = &PortForwardRequest_115{}

func (o *PortForwardRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardRequest{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardRequest)
	}
}
func (o *PortForwardRequest_115) Unwrap() interface{}  { return o.inner }
func (o *PortForwardRequest_115) PodSandboxId() string { return o.inner.PodSandboxId }
func (o *PortForwardRequest_115) SetPodSandboxId(podSandboxId string) {
	o.inner.PodSandboxId = podSandboxId
}

// ---

type PortForwardResponse_115 struct {
	inner *runtimeapi.PortForwardResponse
}

var _ PortForwardResponse = &PortForwardResponse_115{}

func (o *PortForwardResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PortForwardResponse{}
	} else {
		o.inner = v.(*runtimeapi.PortForwardResponse)
	}
}
func (o *PortForwardResponse_115) Unwrap() interface{} { return o.inner }
func (o *PortForwardResponse_115) Url() string         { return o.inner.Url }
func (o *PortForwardResponse_115) SetUrl(url string)   { o.inner.Url = url }

// ---

type ListImagesRequest_115 struct {
	inner *runtimeapi.ListImagesRequest
}

var _ ListImagesRequest = &ListImagesRequest_115{}

func (o *ListImagesRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesRequest{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesRequest)
	}
}
func (o *ListImagesRequest_115) Unwrap() interface{} { return o.inner }
func (o *ListImagesRequest_115) ImageFilter() string { return o.inner.Filter.GetImage().GetImage() }
func (o *ListImagesRequest_115) SetImageFilter(image string) {
	if o.inner.Filter == nil {
		o.inner.Filter = &runtimeapi.ImageFilter{
			Image: &runtimeapi.ImageSpec{Image: image},
		}
	} else {
		o.inner.Filter.Image = &runtimeapi.ImageSpec{Image: image}
	}
}

// ---

type ListImagesResponse_115 struct {
	inner *runtimeapi.ListImagesResponse
}

var _ ListImagesResponse = &ListImagesResponse_115{}

func (o *ListImagesResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ListImagesResponse{}
	} else {
		o.inner = v.(*runtimeapi.ListImagesResponse)
	}
}
func (o *ListImagesResponse_115) Unwrap() interface{} { return o.inner }
func (o *ListImagesResponse_115) Items() []CRIObject {
	var r []CRIObject
	for _, image := range o.inner.Images {
		r = append(r, &Image_115{image})
	}
	return r
}
func (o *ListImagesResponse_115) SetItems(items []CRIObject) {
	o.inner.Images = nil
	for _, wrapped := range items {
		o.inner.Images = append(o.inner.Images, wrapped.Unwrap().(*runtimeapi.Image))
	}
}

// ---

type ImageStatusRequest_115 struct {
	inner *runtimeapi.ImageStatusRequest
}

var _ ImageStatusRequest = &ImageStatusRequest_115{}

func (o *ImageStatusRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusRequest)
	}
}
func (o *ImageStatusRequest_115) Unwrap() interface{} { return o.inner }
func (o *ImageStatusRequest_115) Image() string       { return o.inner.Image.GetImage() }
func (o *ImageStatusRequest_115) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type ImageStatusResponse_115 struct {
	inner *runtimeapi.ImageStatusResponse
}

var _ ImageStatusResponse = &ImageStatusResponse_115{}

func (o *ImageStatusResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageStatusResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageStatusResponse)
	}
}
func (o *ImageStatusResponse_115) Unwrap() interface{} { return o.inner }
func (o *ImageStatusResponse_115) Image() Image {
	if o.inner.Image == nil {
		return nil
	}
	return &Image_115{o.inner.Image}
}
func (o *ImageStatusResponse_115) SetImage(image Image) {
	o.inner.Image = image.Unwrap().(*runtimeapi.Image)
}

// ---

type PullImageRequest_115 struct {
	inner *runtimeapi.PullImageRequest
}

var _ PullImageRequest = &PullImageRequest_115{}

func (o *PullImageRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.PullImageRequest)
	}
}
func (o *PullImageRequest_115) Unwrap() interface{} { return o.inner }
func (o *PullImageRequest_115) Image() string       { return o.inner.Image.GetImage() }
func (o *PullImageRequest_115) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type PullImageResponse_115 struct {
	inner *runtimeapi.PullImageResponse
}

var _ PullImageResponse = &PullImageResponse_115{}

func (o *PullImageResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.PullImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.PullImageResponse)
	}
}
func (o *PullImageResponse_115) Unwrap() interface{}   { return o.inner }
func (o *PullImageResponse_115) Image() string         { return o.inner.ImageRef }
func (o *PullImageResponse_115) SetImage(image string) { o.inner.ImageRef = image }

// ---

type RemoveImageRequest_115 struct {
	inner *runtimeapi.RemoveImageRequest
}

var _ RemoveImageRequest = &RemoveImageRequest_115{}

func (o *RemoveImageRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageRequest{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageRequest)
	}
}
func (o *RemoveImageRequest_115) Unwrap() interface{} { return o.inner }
func (o *RemoveImageRequest_115) Image() string       { return o.inner.Image.GetImage() }
func (o *RemoveImageRequest_115) SetImage(image string) {
	o.inner.Image = &runtimeapi.ImageSpec{Image: image}
}

// ---

type RemoveImageResponse_115 struct {
	inner *runtimeapi.RemoveImageResponse
}

var _ RemoveImageResponse = &RemoveImageResponse_115{}

func (o *RemoveImageResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.RemoveImageResponse{}
	} else {
		o.inner = v.(*runtimeapi.RemoveImageResponse)
	}
}
func (o *RemoveImageResponse_115) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoRequest_115 struct {
	inner *runtimeapi.ImageFsInfoRequest
}

var _ ImageFsInfoRequest = &ImageFsInfoRequest_115{}

func (o *ImageFsInfoRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoRequest{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoRequest)
	}
}
func (o *ImageFsInfoRequest_115) Unwrap() interface{} { return o.inner }

// ---

type ImageFsInfoResponse_115 struct {
	inner *runtimeapi.ImageFsInfoResponse
}

var _ ImageFsInfoResponse = &ImageFsInfoResponse_115{}

func (o *ImageFsInfoResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.ImageFsInfoResponse{}
	} else {
		o.inner = v.(*runtimeapi.ImageFsInfoResponse)
	}
}
func (o *ImageFsInfoResponse_115) Unwrap() interface{} { return o.inner }
func (o *ImageFsInfoResponse_115) Items() []CRIObject {
	var r []CRIObject
	for _, fs := range o.inner.ImageFilesystems {
		r = append(r, &FilesystemUsage_115{fs})
	}
	return r
}
func (o *ImageFsInfoResponse_115) SetItems(items []CRIObject) {
	o.inner.ImageFilesystems = nil
	for _, wrapped := range items {
		o.inner.ImageFilesystems = append(o.inner.ImageFilesystems, wrapped.Unwrap().(*runtimeapi.FilesystemUsage))
	}
}

// --- 1.8+ only ---

type UpdateContainerResourcesRequest_115 struct {
	inner *runtimeapi.UpdateContainerResourcesRequest
}

var _ UpdateContainerResourcesRequest = &UpdateContainerResourcesRequest_115{}

func (o *UpdateContainerResourcesRequest_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesRequest{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesRequest)
	}
}
func (o *UpdateContainerResourcesRequest_115) Unwrap() interface{}      { return o.inner }
func (o *UpdateContainerResourcesRequest_115) ContainerId() string      { return o.inner.ContainerId }
func (o *UpdateContainerResourcesRequest_115) SetContainerId(id string) { o.inner.ContainerId = id }

// --- 1.8+ only ---

type UpdateContainerResourcesResponse_115 struct {
	inner *runtimeapi.UpdateContainerResourcesResponse
}

var _ UpdateContainerResourcesResponse = &UpdateContainerResourcesResponse_115{}

func (o *UpdateContainerResourcesResponse_115) Wrap(v interface{}) {
	if v == nil {
		o.inner = &runtimeapi.UpdateContainerResourcesResponse{}
	} else {
		o.inner = v.(*runtimeapi.UpdateContainerResourcesResponse)
	}
}
func (o *UpdateContainerResourcesResponse_115) Unwrap() interface{} { return o.inner }

// ---

var cri115typeMatcher = newTypeMatcher()

func init() {
	cri115typeMatcher.registerTypes(
		&PodSandbox_115{},
		&Container_115{},
		&Image_115{},
		&PodSandboxStatus_115{},
		&ContainerStatus_115{},
		&ContainerStats_115{},
		&FilesystemUsage_115{},
		&VersionRequest_115{},
		&VersionResponse_115{},
		&StatusRequest_115{},
		&StatusResponse_115{},
		&UpdateRuntimeConfigRequest_115{},
		&UpdateRuntimeConfigResponse_115{},
		&RunPodSandboxRequest_115{},
		&RunPodSandboxResponse_115{},
		&ListPodSandboxRequest_115{},
		&ListPodSandboxResponse_115{},
		&StopPodSandboxRequest_115{},
		&StopPodSandboxResponse_115{},
		&RemovePodSandboxRequest_115{},
		&RemovePodSandboxResponse_115{},
		&PodSandboxStatusRequest_115{},
		&PodSandboxStatusResponse_115{},
		&CreateContainerRequest_115{},
		&CreateContainerResponse_115{},
		&ListContainersRequest_115{},
		&ListContainersResponse_115{},
		&ListContainerStatsRequest_115{},
		&ListContainerStatsResponse_115{},
		&StartContainerRequest_115{},
		&StartContainerResponse_115{},
		&StopContainerRequest_115{},
		&StopContainerResponse_115{},
		&RemoveContainerRequest_115{},
		&RemoveContainerResponse_115{},
		&ReopenContainerLogRequest_115{},
		&ReopenContainerLogResponse_115{},
		&ContainerStatusRequest_115{},
		&ContainerStatusResponse_115{},
		&ContainerStatsRequest_115{},
		&ContainerStatsResponse_115{},
		&ExecSyncRequest_115{},
		&ExecSyncResponse_115{},
		&ExecRequest_115{},
		&ExecResponse_115{},
		&AttachRequest_115{},
		&AttachResponse_115{},
		&PortForwardRequest_115{},
		&PortForwardResponse_115{},
		&ListImagesRequest_115{},
		&ListImagesResponse_115{},
		&ImageStatusRequest_115{},
		&ImageStatusResponse_115{},
		&PullImageRequest_115{},
		&PullImageResponse_115{},
		&RemoveImageRequest_115{},
		&RemoveImageResponse_115{},
		&ImageFsInfoRequest_115{},
		&ImageFsInfoResponse_115{},
		&UpdateContainerResourcesRequest_115{},
		&UpdateContainerResourcesResponse_115{},
	)
}

// CRI115 denotes the CRI version 1.10
type CRI115 struct{}

var _ CRIVersion = &CRI115{}

func (c *CRI115) Register(server *grpc.Server) {
	runtimeapi.RegisterDummyRuntimeServiceServer(server)
	runtimeapi.RegisterDummyImageServiceServer(server)
}

func (c *CRI115) ProbeRequest() (interface{}, interface{}) {
	return &runtimeapi.VersionRequest{}, &runtimeapi.VersionResponse{}
}

func (c *CRI115) WrapObject(o interface{}) (CRIObject, CRIObject, error) {
	return wrapUsingMatcher(cri115typeMatcher, o)
}

func (c *CRI115) ProtoPackage() string { return "runtime.v1alpha2" }

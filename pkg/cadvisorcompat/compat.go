// Copyright 2023 Francesco Romani (fromani at gmail.com)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cadvisorcompat

import (
	"fmt"

	"github.com/google/cadvisor/events"
	cadvisorapi "github.com/google/cadvisor/info/v1"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"

	"github.com/ffromani/ghwadvisor/pkg/ghwext/machine"
	"github.com/jaypipes/ghw/pkg/option"
)

var (
	ErrNotImplemented = fmt.Errorf("not implemented")
)

type Handle struct {
	sysfsRoot string
	info      *machine.Info
}

func NewFactory(sysfsRoot string) (*Handle, error) {
	hnd := New(sysfsRoot)
	err := hnd.Start()
	return hnd, err
}

func New(sysfsRoot string) *Handle {
	return &Handle{
		sysfsRoot: sysfsRoot,
	}
}

func (hnd *Handle) Start() error {
	pathOverrides := option.PathOverrides{
		"/sys": hnd.sysfsRoot,
	}
	info, err := machine.New(
		option.WithDisableTools(),
		option.WithPathOverrides(pathOverrides),
	)
	hnd.info = info
	return err
}

func (hnd *Handle) DockerContainer(name string, req *cadvisorapi.ContainerInfoRequest) (cadvisorapi.ContainerInfo, error) {
	return cadvisorapi.ContainerInfo{}, ErrNotImplemented
}

func (hnd *Handle) ContainerInfo(name string, req *cadvisorapi.ContainerInfoRequest) (*cadvisorapi.ContainerInfo, error) {
	return nil, ErrNotImplemented
}

func (hnd *Handle) ContainerInfoV2(name string, options cadvisorapiv2.RequestOptions) (map[string]cadvisorapiv2.ContainerInfo, error) {
	return nil, ErrNotImplemented
}

func (hnd *Handle) GetRequestedContainersInfo(containerName string, options cadvisorapiv2.RequestOptions) (map[string]*cadvisorapi.ContainerInfo, error) {
	return nil, ErrNotImplemented
}

func (hnd *Handle) SubcontainerInfo(name string, req *cadvisorapi.ContainerInfoRequest) (map[string]*cadvisorapi.ContainerInfo, error) {
	return nil, ErrNotImplemented
}

func (hnd *Handle) MachineInfo() (*cadvisorapi.MachineInfo, error) {
	return hnd.info.ToCAdvisorMachineInfo(), nil
}

func (hhd Handle) VersionInfo() (*cadvisorapi.VersionInfo, error) {
	return nil, nil // TODO
}

func (hnd *Handle) ImagesFsInfo() (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, ErrNotImplemented
}

func (hnd *Handle) RootFsInfo() (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, ErrNotImplemented
}

func (hnd *Handle) WatchEvents(request *events.Request) (*events.EventChannel, error) {
	return nil, ErrNotImplemented
}

func (hnd *Handle) GetDirFsInfo(path string) (cadvisorapiv2.FsInfo, error) {
	return cadvisorapiv2.FsInfo{}, nil
}

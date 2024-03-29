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

package machine

import (
	"github.com/jaypipes/ghw/pkg/block"
	"github.com/jaypipes/ghw/pkg/memory"
	"github.com/jaypipes/ghw/pkg/net"
	"github.com/jaypipes/ghw/pkg/option"
	"github.com/jaypipes/ghw/pkg/topology"

	"github.com/ffromani/ghwadvisor/pkg/ghwext/cpu"
)

type Info struct {
	Memory   *memory.Info   `json:"memory"`
	Block    *block.Info    `json:"block"`
	CPU      *cpu.Info      `json:"cpu"`
	Topology *topology.Info `json:"topology"`
	Network  *net.Info      `json:"network"`
	// future: note: pcidb load is hanging on wasmedge 0.13.3
	// when compiled with golang 1.21.{0,1}
}

func New(opts ...*option.Option) (*Info, error) {
	memInfo, err := memory.New(opts...)
	if err != nil {
		return nil, err
	}
	blockInfo, err := block.New(opts...)
	if err != nil {
		return nil, err
	}
	cpuInfo, err := cpu.New(opts...)
	if err != nil {
		return nil, err
	}
	topologyInfo, err := topology.New(opts...)
	if err != nil {
		return nil, err
	}
	netInfo, err := net.New(opts...)
	if err != nil {
		return nil, err
	}
	return &Info{
		CPU:      cpuInfo,
		Memory:   memInfo,
		Block:    blockInfo,
		Topology: topologyInfo,
		Network:  netInfo,
	}, nil
}

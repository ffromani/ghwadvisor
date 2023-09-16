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

package cpu

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/jaypipes/ghw/pkg/context"
	ghwcpu "github.com/jaypipes/ghw/pkg/cpu"
	"github.com/jaypipes/ghw/pkg/linuxpath"
	"github.com/jaypipes/ghw/pkg/option"

	"k8s.io/klog/v2"
)

type Info struct {
	ghwcpu.Info
	// TotalPackages is the total number of physical packages the host system
	// contains
	TotalPackages uint32 `json:"total_packages"`
	// CurrentFrequency is the processor Frequency in KHz when the information was
	// gathered. On multi-core and multi-package system this value is meaningless.
	CurrentFrequency uint32 `json:"current_frequency"`
}

func New(opts ...*option.Option) (*Info, error) {
	base, err := ghwcpu.New(opts...)
	if err != nil {
		return nil, err
	}
	ctx := context.New(opts...)
	return &Info{
		Info:             *base,
		TotalPackages:    uint32(NumPhysicalPackages(base)),
		CurrentFrequency: uint32(CurrentFrequencyKHz(ctx)),
	}, nil
}

// "The kernel does not care about the concept of physical sockets because a socket
// has no relevance to software. It’s an electromechanical component.
// In the past a socket always contained a single package (see below),
// but with the advent of Multi Chip Modules (MCM) a socket can hold more than one
// package. So there might be still references to sockets in the code, but they are
// of historical nature and should be cleaned up."
// -- https://www.kernel.org/doc/html/v5.9/x86/topology.html#x86-topology
func NumPhysicalPackages(cpuInfo *ghwcpu.Info) int {
	cpus := make(map[int]struct{})
	for _, proc := range cpuInfo.Processors {
		cpus[proc.ID] = struct{}{}
	}
	return len(cpus)
}

func CurrentFrequencyKHz(ctx *context.Context) int {
	paths := linuxpath.New(ctx)
	// hardcoded cpuID because we must always have cpu0ø
	curFreqPath := filepath.Join(paths.SysDevicesSystemCPU, "cpu0", "cpufreq", "scaling_cur_freq")
	curFreqData, err := os.ReadFile(curFreqPath)
	if err != nil {
		klog.Warning("cannot read frequency data from %q: %v", curFreqPath, err)
		return 0
	}
	curFreq, err := strconv.Atoi(strings.TrimSpace(string(curFreqData)))
	if err != nil {
		klog.Warning("cannot parse frequency data (%q) from %q: %v", curFreqData, curFreqPath, err)
		return 0
	}
	return curFreq
}

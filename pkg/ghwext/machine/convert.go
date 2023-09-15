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
	"time"

	cadvisorapiv1 "github.com/google/cadvisor/info/v1"
)

func (info *Info) IntoCAdvisorMachineInfo(machineInfo *cadvisorapiv1.MachineInfo) {
	machineInfo.CPUVendorID = info.CPU.Processors[0].Vendor
	machineInfo.NumCores = int(info.CPU.TotalCores)
	machineInfo.NumPhysicalCores = int(info.CPU.TotalThreads)
	machineInfo.NumSockets = 0   // TODO
	machineInfo.CpuFrequency = 0 // TODO

	machineInfo.MemoryByType = make(map[string]*cadvisorapiv1.MemoryInfo)
	machineInfo.HugePages = []cadvisorapiv1.HugePagesInfo{} // TODO

	machineInfo.Filesystems = []cadvisorapiv1.FsInfo{} // TODO: maybe

	machineInfo.DiskMap = make(map[string]cadvisorapiv1.DiskInfo)

	machineInfo.NetworkDevices = []cadvisorapiv1.NetInfo{}

	machineInfo.Topology = []cadvisorapiv1.Node{}

	machineInfo.BootID = ""     // explicitly none
	machineInfo.MachineID = ""  // explicitly none
	machineInfo.SystemUUID = "" // explicitly none

	machineInfo.CloudProvider = "Unknown"
	machineInfo.InstanceType = "Unknown"
	machineInfo.InstanceID = "None"

	machineInfo.Timestamp = time.Now()
}
func (info *Info) ToCAdvisorMachineInfo() *cadvisorapiv1.MachineInfo {
	var machineInfo cadvisorapiv1.MachineInfo
	info.IntoCAdvisorMachineInfo(&machineInfo)
	return &machineInfo
}

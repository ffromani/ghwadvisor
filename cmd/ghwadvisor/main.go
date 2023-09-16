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

package main

import (
	"encoding/json"
	goflag "flag"
	"fmt"
	"log"
	"net/http"
	"time"

	flag "github.com/spf13/pflag"

	cadvisorapiv1 "github.com/google/cadvisor/info/v1"

	"k8s.io/klog/v2"

	"github.com/ffromani/ghwadvisor/pkg/cadvisorcompat"
	"github.com/ffromani/ghwadvisor/pkg/router"
)

func main() {
	klog.InitFlags(nil)

	var port int
	var sysfsRoot string
	flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	flag.IntVar(&port, "port", 8080, "port to listen")
	flag.StringVar(&sysfsRoot, "sysfs", "/sys", "sysfs mount point - use when run inside a container")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)

	klog.Infof("ghwadvisor starting on [%s]", addr)
	defer klog.Infof("ghwadvisor stopped")

	setupStart := time.Now()

	klog.Infof("gathering machine data from %q", sysfsRoot)
	intf, err := cadvisorcompat.NewFactory(sysfsRoot)
	if err != nil {
		klog.Fatalf("failed to collect system data: %v", err)
	}

	minfo, err := intf.MachineInfo()
	if err != nil {
		klog.Fatalf("failed to collect machine info: %v", err)
	}

	mh := machineHandler{machineInfo: minfo}
	rt := router.New([]router.Route{
		{
			"machine",
			"GET",
			"/api/v1.3/machine", // see: https://github.com/google/cadvisor/blob/master/docs/api.md
			mh.ServeHTTP,
		},
	})

	klog.Infof("ghwadvisor ready (setup time: %v)", time.Since(setupStart))

	klog.Infof("machine data:\n%s", toJSON(mh.machineInfo))
	log.Fatal(http.ListenAndServe(addr, rt))
}

type machineHandler struct {
	machineInfo *cadvisorapiv1.MachineInfo
}

func (mh machineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(mh.machineInfo); err != nil {
		panic(err)
	}
}

func toJSON(obj any) string {
	data, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Sprintf("<ERROR: %v>", err)
	}
	return string(data)
}

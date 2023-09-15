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
	"fmt"
	"log"
	"net/http"
	"time"

	cadvisorapiv1 "github.com/google/cadvisor/info/v1"

	"k8s.io/klog/v2"

	"github.com/ffromani/ghwadvisor/pkg/ghwext/machine"
	"github.com/ffromani/ghwadvisor/pkg/router"
)

func main() {
	klog.Infof("ghwadvisor starting")
	defer klog.Infof("ghwadvisor stopped")

	setupStart := time.Now()

	info, err := machine.New()
	if err != nil {
		klog.Fatalf("ghw failed: %v", err)
	}

	mh := machineHandler{machineInfo: info.ToCAdvisorMachineInfo()}
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
	log.Fatal(http.ListenAndServe(":8080", rt))
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

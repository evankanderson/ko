// Copyright 2018 Google LLC All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

func main() {
	dp := os.Getenv("KO_DATA_PATH")
	file := filepath.Join(dp, "kenobi")
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading %q: %v", file, err)
	}
	log.Print(string(bytes))

	file = filepath.Join(dp, "refs/heads/master")
	bytes, err = ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading %q: %v", file, err)
	}
	log.Printf(string(bytes))

	// Cause the pod to "hang" to allow us to check for a readiness state.
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}
